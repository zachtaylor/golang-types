package types

import "net/http"

// HTTPPather is Handler and Router
type HTTPPather interface {
	HTTPRouter
	HTTPServer
}

// HTTPRequest is http.Request
type HTTPRequest = http.Request

// HTTPRouter is an HTTP routing interface
type HTTPRouter interface {
	RouteHTTP(*http.Request) bool
}

// HTTPServer is http.Handler
type HTTPServer = http.Handler

// HTTPServerFunc is http.HandlerFunc
type HTTPServerFunc = http.HandlerFunc

// HTTPWriter is http.ResponseWriter
type HTTPWriter = http.ResponseWriter
