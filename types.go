package avicenna

import "net/http"

type Router struct {
	NotFoundHandler         http.Handler
	MethodNotAllowedHandler http.Handler
	Routes                  map[Route]http.Handler
}

type Route struct {
	Method string
	Path   string
}
