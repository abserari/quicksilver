// from https://www.integralist.co.uk/posts/golang-reverse-proxy/
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var originhost = "http://localhost:9000/"

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "backend server handled the request!")
	}))
	defer backendServer.Close()

	backendServerURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = backendServerURL.Scheme
			r.URL.Host = backendServerURL.Host
			r.URL.Path = singleJoiningSlash(backendServerURL.Path, r.URL.Path)
		},
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		},
		ModifyResponse: func(r *http.Response) error {
			// return nil
			//
			// purposefully return an error so ErrorHandler gets called
			return errors.New("uh-oh")
		},
		ErrorHandler: func(rw http.ResponseWriter, r *http.Request, err error) {
			fmt.Printf("error was: %+v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		},
	}

	frontendServer := httptest.NewServer(proxy)
	defer frontendServer.Close()

	resp, err := http.Get(frontendServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nbody: \n%s\n\n", b)
}
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
