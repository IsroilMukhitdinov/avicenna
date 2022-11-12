package avicenna

import (
	"net/http"
)

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path

	route := Route{
		Method: method,
		Path:   path,
	}

	if !router.isValidMethod(method) || !router.isValidPath(path) || !router.isValidRoute(route) {
		router.NotFoundHandler.ServeHTTP(w, r)
		return
	}

	handler := router.Routes[route]
	handler.ServeHTTP(w, r)
}

func NewRouter(routes map[Route]http.Handler) *Router {
	newRouter := &Router{
		NotFoundHandler:         defaultNotFoundHandler,
		MethodNotAllowedHandler: defaultMethodNotAllowedHandler,
	}

	return newRouter
}
