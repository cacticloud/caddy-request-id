package caddyrequestid

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/cacticloud/caddy-request-id"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(RequestID{})
}

// RequestID is a Caddy module that adds a unique request ID to each request.
type RequestID struct{}

// CaddyModule returns the Caddy module information.
func (RequestID) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.request_id",
		New: func() caddy.Module { return new(RequestID) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (r RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Generate a unique ID based on the current time and a random component.
	uniqueID := generateUniqueID()
	// Set the X-Request-ID header with the unique ID.
	r.Header.Set("X-Request-ID", uniqueID)

	// Call the next handler in the chain.
	return next.ServeHTTP(w, r)
}

// generateUniqueID creates a unique ID using a combination of time and a random number.
func generateUniqueID() string {
	now := time.Now()
	return fmt.Sprintf("%d%06d", now.UnixNano(), rand.Intn(999999))
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*RequestID)(nil)
	_ caddy.Provisioner           = (*RequestID)(nil)
)
