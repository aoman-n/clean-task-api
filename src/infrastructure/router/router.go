package router

import (
	"task-api/src/infrastructure/middleware"
	"task-api/src/interfaces"
	"task-api/src/usecase"
)

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *Engine {

	middlewares := middleware.New(sqlhandler)

	userController := interfaces.NewUserController(sqlhandler, validator)
	projectController := interfaces.NewProjectController(sqlhandler, validator)
	taskController := interfaces.NewTaskController(sqlhandler, validator)
	tagController := interfaces.NewTagController(sqlhandler, validator)

	router := NewH()

	/* users API */
	router.POST("/signup", userController.Singup)
	router.POST("/login", userController.Login)
	router.GET("/users", userController.Index)
	router.GET("/users/:id", userController.Show)

	// /* projects API */
	router.GET("/projects", middlewares.Authenticate, projectController.Index)
	router.POST("/projects", middlewares.Authenticate, projectController.Create)
	router.DELETE("/projects/:id", middlewares.Authenticate, projectController.Delete)

	// /* task API */
	router.GET("/projects/:id/tasks", middlewares.Authenticate, middlewares.RequiredJoinedProject, taskController.Index)
	router.POST("/projects/:id/tasks", middlewares.Authenticate, middlewares.RequiredWriteRole, taskController.Create)
	router.DELETE("/projects/:id/tasks/:task_id", middlewares.Authenticate, middlewares.RequiredWriteRole, taskController.Delete)
	router.PUT("/projects/:id/tasks/:task_id", middlewares.Authenticate, middlewares.RequiredWriteRole, taskController.Update)

	// /* tags API */
	router.GET("/projects/:id/tags", middlewares.Authenticate, middlewares.RequiredWriteRole, tagController.Index)
	router.POST("/projects/:id/tags", middlewares.Authenticate, middlewares.RequiredWriteRole, tagController.Create)
	router.PUT("/tags/:id", middlewares.Authenticate, tagController.Update)
	router.DELETE("/tags/:id", middlewares.Authenticate, tagController.Delete)

	return router
}
