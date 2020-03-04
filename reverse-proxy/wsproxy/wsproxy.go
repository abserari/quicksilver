package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if upgradeType(r.Header) != "websocket" {
			log.Error("unexpected backend request")
			http.Error(w, "unexpected request", 400)
			return
		}
		c, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			log.Error(err)
			return
		}
		defer c.Close()
		io.WriteString(c, "HTTP/1.1 101 Switching Protocols\r\nConnection: upgrade\r\nUpgrade: WebSocket\r\n\r\n")
		bs := bufio.NewScanner(c)
		if !bs.Scan() {
			log.Errorf("backend failed to read line from client: %v", bs.Err())
			return
		}
		fmt.Fprintf(c, "backend got %q\n", bs.Text())
	}))
	defer backendServer.Close()

	backURL, _ := url.Parse(backendServer.URL)
	rproxy := NewSingleHostReverseProxy(backURL)
	rproxy.ErrorLog = log.New(ioutil.Discard, "", 0) // quiet for tests

	frontendProxy := httptest.NewServer(rproxy)
	defer frontendProxy.Close()

	req, _ := http.NewRequest("GET", frontendProxy.URL, nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")

	c := frontendProxy.Client()
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 101 {
		log.Fatalf("status = %v; want 101", res.Status)
	}
	if upgradeType(res.Header) != "websocket" {
		log.Fatalf("not websocket upgrade; got %#v", res.Header)
	}
	rwc, ok := res.Body.(io.ReadWriteCloser)
	if !ok {
		log.Fatalf("response body is of type %T; does not implement ReadWriteCloser", res.Body)
	}
	defer rwc.Close()

	io.WriteString(rwc, "Hello\n")
	bs := bufio.NewScanner(rwc)
	if !bs.Scan() {
		log.Fatalf("Scan: %v", bs.Err())
	}
	got := bs.Text()
	want := `backend got "Hello"`
	if got != want {
		log.Errorf("got %#q, want %#q", got, want)
	}
}
