package engine

import (
	"fmt"
	"net/http"
	"task-api/src/interfaces"

	"github.com/julienschmidt/httprouter"
)

func log(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ACCESS] ", r.Method, r.URL, r.Host, r.RequestURI)
}

type Handle interface {
	Dispatch(handlers []HandlerFunc) httprouter.Handle
}

type handle struct{}

func (h *handle) Dispatch(handlers []HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log(w, r)
		ctx := Context{w, r, ps, true}
		for _, h := range handlers {
			if ctx.enable {
				h(&ctx)
			}
		}
	}
}

type Engine struct {
	router    *httprouter.Router
	handle    Handle
	groupPath string
}

func New() *Engine {
	router := httprouter.New()
	handle := &handle{}

	return &Engine{router, handle, ""}
}

type HandlerFunc func(c interfaces.Context)

func (e *Engine) Group(path string) {
	e.groupPath = path
}

func (e *Engine) GET(path string, handlers ...HandlerFunc) {
	e.router.GET(e.groupPath+path, e.handle.Dispatch(handlers))
}

func (e *Engine) POST(path string, handlers ...HandlerFunc) {
	e.router.POST(e.groupPath+path, e.handle.Dispatch(handlers))
}

func (e *Engine) PUT(path string, handlers ...HandlerFunc) {
	e.router.PUT(e.groupPath+path, e.handle.Dispatch(handlers))
}

func (e *Engine) DELETE(path string, handlers ...HandlerFunc) {
	e.router.DELETE(e.groupPath+path, e.handle.Dispatch(handlers))
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.ServeHTTP(w, r)
}
