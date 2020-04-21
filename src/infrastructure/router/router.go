package router

import (
	"fmt"
	"net/http"
	"task-api/src/interfaces"

	"github.com/julienschmidt/httprouter"
)

func logging(h interfaces.HttpHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("[ACCESS] ", r.Method, r.URL, r.Host, r.RequestURI)
		h(w, r, params)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World"))
}

func Handler(sqlhandler interfaces.SQLHandler) *httprouter.Router {
	userController := interfaces.NewUserController(sqlhandler)

	router := httprouter.New()
	router.GET("/users", logging(userController.Index))

	return router
}
