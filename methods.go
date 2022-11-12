package avicenna

import (
	"net/http"
)

func (router *Router) Method(method string, path string, handlerFunc http.HandlerFunc) {
	if !router.isValidPath(path) {
		return
	}
	route := Route{
		Method: method,
		Path:   path,
	}

	router.Routes[route] = http.HandlerFunc(handlerFunc)
}

func (router *Router) Get(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodGet, path, handlerFunc)
}

func (router *Router) Post(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodPost, path, handlerFunc)
}

func (router *Router) Delete(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodDelete, path, handlerFunc)
}

func (router *Router) Put(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodPut, path, handlerFunc)
}

func (router *Router) Patch(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodPatch, path, handlerFunc)
}

func (router *Router) Connect(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodConnect, path, handlerFunc)
}

func (router *Router) Head(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodHead, path, handlerFunc)
}

func (router *Router) Trace(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodTrace, path, handlerFunc)
}

func (router *Router) Options(path string, handlerFunc http.HandlerFunc) {
	router.Method(http.MethodOptions, path, handlerFunc)
}
