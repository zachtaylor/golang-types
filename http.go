package types

import "net/http"

// HTTPHandler is http.Handler
type HTTPHandler = http.Handler

// HTTPPather is Handler and Router
type HTTPPather interface {
	HTTPHandler
	HTTPRouter
}

// HTTPRequest is http.Request
type HTTPRequest = http.Request

// HTTPRouter is an HTTP routing interface
type HTTPRouter interface {
	RouteHTTP(*http.Request) bool
}

// HTTPWriter is http.ResponseWriter
type HTTPWriter = http.ResponseWriter
