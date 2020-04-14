package router

import (
	"fmt"
	"net/http"
	"task-api/src/infrastructure/db"
	"task-api/src/interfaces"

	"github.com/julienschmidt/httprouter"
)

func logging(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("recieved access: ", r.Method, r.URL, r.Host, r.RequestURI)
		h(w, r, params)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World"))
}

func Handler(sqlhandler *db.Sqlhandler) *httprouter.Router {
	userController := interfaces.NewUserController(sqlhandler)

	router := httprouter.New()
	router.GET("/", logging(userController.Index))

	return router
}
