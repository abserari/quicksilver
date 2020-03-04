package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
)

type override struct {
	Header string
	Match  string
	Host   string
	Path   string
}

type config struct {
	Path     string
	Host     string
	Override override
}

var configuration = []config{
	config{
		Path: "/api/room",
		Host: "httpbin.org",
	},
	config{
		Path: "/api/foo",
		Host: "httpbin.org",
	},
	config{
		Path: "/api/student",
		Host: "httpbin.org",
	},
}

func main() {
	r := mux.NewRouter()

	for _, conf := range configuration {
		proxy := generateProxy(conf)

		r.HandleFunc(conf.Path, func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	}

	log.Fatal(http.ListenAndServe(":9001", r))
}

func generateProxy(conf config) http.Handler {
	proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
		originHost := conf.Host
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", originHost)
		req.Host = originHost
		req.URL.Host = originHost
		req.URL.Scheme = "https"

		if conf.Override.Header != "" && conf.Override.Match != "" {
			if req.Header.Get(conf.Override.Header) == conf.Override.Match {
				req.URL.Path = conf.Override.Path
			}
		}
	}, Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
	}}

	return proxy
}
