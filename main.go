package main

import (
	"log"
	"net/http"
	"task-api/src/config"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World"))
}

func main() {
	conf := config.Get()

	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":"+conf.Server.Port, router))
}
