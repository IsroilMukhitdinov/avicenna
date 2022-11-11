package avicenna

import (
	"net/http"

	"github.com/IsroilMukhitdionov/avicenna/internal"
)

var defaultNotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
})

type Router struct {
	NotFoundHandler http.Handler
	Routes          map[Route]http.Handler
}

type Route struct {
	Method string
	Path   string
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func NewRouter(notFoundHandler http.Handler, routes map[Route]http.Handler) *Router {
	newRouter := &Router{
		Routes: routes,
	}

	if notFoundHandler == nil {
		newRouter.NotFoundHandler = defaultNotFoundHandler
	} else {
		newRouter.NotFoundHandler = notFoundHandler
	}

	return newRouter
}

func (router *Router) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	if !internal.IsValidPath(path) {
		return
	}

	route := Route{
		Method: "GET",
		Path:   path,
	}

	router.Routes[route] = http.HandlerFunc(handler)
}
