package types

import "net/http"

// HTTPPath is a simple HTTPPather
type HTTPPath struct {
	Router HTTPRouter
	Server HTTPServer
}

// RouteHTTP implements `HTTPRouter` by calling calling the internal `HTTPRouter`
func (p HTTPPath) RouteHTTP(r *HTTPRequest) bool { return p.Router.RouteHTTP(r) }

// ServeHTTP implements `HTTPServer` by calling calling the internal `HTTPServer`
func (p HTTPPath) ServeHTTP(w HTTPWriter, r *HTTPRequest) { p.Server.ServeHTTP(w, r) }

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

// HTTPTree is a HTTPServer made of []HTTPPather
type HTTPTree []HTTPPather

// Add calls Branch with a new HTTPPath
func (t *HTTPTree) Add(r HTTPRouter, s HTTPServer) {
	t.Branch(HTTPPath{
		Router: r,
		Server: s,
	})
}

// Branch appends a HTTPPather to this HTTPTree
func (t *HTTPTree) Branch(p HTTPPather) { *t = append(*t, p) }

// ServeHTTP implements HTTPServer by pathing to a branch
func (t *HTTPTree) ServeHTTP(w HTTPWriter, r *HTTPRequest) {
	var s HTTPServer
	for _, p := range *t {
		if p.RouteHTTP(r) {
			s = p
			break
		}
	}
	if s != nil {
		s.ServeHTTP(w, r)
	}
}

// HTTPWriter is http.ResponseWriter
type HTTPWriter = http.ResponseWriter
