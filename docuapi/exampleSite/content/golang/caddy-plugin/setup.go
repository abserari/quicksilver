package gizmo

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("something", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}
func setup(c *caddy.Controller) error {
	g := gizmoHandler{}
	for c.Next() {
		if !c.NextArg() {
			return c.ArgErr()
		}
		_ = c.Val()

		httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
			g.next = next
			return g
		})
	}
	return nil
}
