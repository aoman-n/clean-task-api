package router

import (
	"task-api/src/interfaces"
	"task-api/src/interfaces/middleware"
	"task-api/src/usecase"
)

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *Engine {

	middleware := middleware.NewMiddlewre(sqlhandler)

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

	return router
}
