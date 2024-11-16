package routers

import (
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	mux *http.ServeMux
}

func (r *Router) Handler(api string, handler handlerFunc) {
	r.mux.Handle(api, http.HandlerFunc(handler))
}

func GetRouter(mux *http.ServeMux) *Router {
	Router := &Router{
		mux: mux,
	}

	return Router
}
