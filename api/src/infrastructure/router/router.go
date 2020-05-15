package router

import (
	"net/http"
	"os"
	"task-api/src/infrastructure/engine"
	"task-api/src/interfaces"
	"task-api/src/interfaces/controller"
	"task-api/src/interfaces/middleware"
	"task-api/src/usecase"

	"github.com/rs/cors"
)

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) http.Handler {

	middleware := middleware.NewMiddlewre(sqlhandler)

	userController := controller.NewUserController(sqlhandler, validator)
	projectController := controller.NewProjectController(sqlhandler, validator)
	taskController := controller.NewTaskController(sqlhandler, validator)
	tagController := controller.NewTagController(sqlhandler, validator)

	router := engine.New()
	router.Group("/")

	/* users API */
	router.POST("/signup", userController.Singup)
	router.POST("/login", userController.Login)
	router.GET("/users", userController.Index)
	router.GET("/users/:id", userController.Show)

	// /* projects API */
	router.GET("/projects", middleware.Auth, projectController.Index)
	router.POST("/projects", middleware.Auth, projectController.Create)
	router.DELETE("/projects/:id", middleware.Auth, projectController.Delete)
	// TODO: show project

	// /* task API */
	router.GET("/projects/:id/tasks", middleware.Auth, middleware.RequiredJoinedProject, taskController.Index)
	router.POST("/projects/:id/tasks", middleware.Auth, middleware.RequiredWriteRole, taskController.Create)
	router.DELETE("/projects/:id/tasks/:task_id", middleware.Auth, middleware.RequiredWriteRole, taskController.Delete)
	router.PUT("/projects/:id/tasks/:task_id", middleware.Auth, middleware.RequiredWriteRole, taskController.Update)

	// /* tags API */
	router.GET("/projects/:id/tags", middleware.Auth, middleware.RequiredWriteRole, tagController.Index)
	router.POST("/projects/:id/tags", middleware.Auth, middleware.RequiredWriteRole, tagController.Create)
	router.PUT("/tags/:id", middleware.Auth, tagController.Update)
	router.DELETE("/tags/:id", middleware.Auth, tagController.Delete)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONT_ORIGIN"), os.Getenv("BFF_ORIGIN")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(router)

	return handler
}
