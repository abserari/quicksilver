package gizmo

import (
	"net/http"

	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

type gizmoHandler struct {
	next httpserver.Handler
}

func (g gizmoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	return g.next.ServeHTTP(w, r)
}
