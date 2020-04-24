package router

import (
	"fmt"
	"net/http"
	"task-api/src/infrastructure/middleware"
	"task-api/src/interfaces"
	"task-api/src/usecase"

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

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *httprouter.Router {
	middlewares := middleware.New(sqlhandler)

	userController := interfaces.NewUserController(sqlhandler, validator)
	projectController := interfaces.NewProjectController(sqlhandler, validator)

	router := httprouter.New()
	/* users API */
	router.POST("/signup", logging(userController.Singup))
	router.POST("/login", logging(userController.Login))
	router.GET("/users", logging(userController.Index))
	router.GET("/users/:id", logging(userController.Show))

	/* projects API */
	router.GET("/projects", logging(middlewares.Authenticate(projectController.Index)))
	router.POST("/projects", logging(middlewares.Authenticate(projectController.Create)))

	return router
}
