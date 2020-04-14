package infrastructure

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World"))
}

func Handler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)

	return router
}
