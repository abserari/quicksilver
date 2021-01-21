package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"

	"golang.org/x/net/http/httpguts"
)

func upgradeType(h http.Header) string {
	if !httpguts.HeaderValuesContainsToken(h["Connection"], "Upgrade") {
		return ""
	}
	return strings.ToLower(h.Get("Upgrade"))
}

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if upgradeType(r.Header) != "websocket" {
			log.Println("unexpected backend request")
			http.Error(w, "unexpected request", 400)
			return
		}
		log.Println("backend server get the websocket connection")
		c, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			log.Println(err)
			return
		}
		defer c.Close()
		log.Println("backend server upgrade http/1.1 101 to websocket")
		io.WriteString(c, "HTTP/1.1 101 Switching Protocols\r\nConnection: upgrade\r\nUpgrade: WebSocket\r\n\r\n")
		bs := bufio.NewScanner(c)
		if !bs.Scan() {
			log.Println(fmt.Errorf("backend failed to read line from client: %v", bs.Err()))
			return
		}
		fmt.Fprintf(c, "backend got %q\n", bs.Text())
	}))
	defer backendServer.Close()

	backURL, _ := url.Parse(backendServer.URL)
	rproxy := httputil.NewSingleHostReverseProxy(backURL)
	rproxy.ErrorLog = log.New(ioutil.Discard, "", 0) // quiet for tests

	frontendProxy := httptest.NewServer(rproxy)
	defer frontendProxy.Close()

	// do request to frontendProxy and forward it to backend server
	req, _ := http.NewRequest("GET", frontendProxy.URL, nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")

	// new client to frontendProxy to do request
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
	log.Println("frontproxy got ws ReadWriteCloser rwc")
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
	log.Println("frontproxy read from response and got:", got)
	want := `backend got "Hello"`
	if got != want {
		log.Println(fmt.Errorf("got %#q, want %#q", got, want))
	}
}
