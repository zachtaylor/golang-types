package types

import "net/http"

// Handler is http.Handler
type Handler = http.Handler

// Pather is Handler and Router
type Pather interface {
	Handler
	Router
}

// Router is an HTTP routing interface
type Router interface {
	RouteHTTP(*http.Request) bool
}
