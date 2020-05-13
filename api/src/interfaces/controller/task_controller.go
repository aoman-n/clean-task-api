package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"task-api/src/interfaces"
	"task-api/src/interfaces/gateway"
	"task-api/src/usecase"
	"task-api/src/usecase/interactor"
	"task-api/src/utils/errors"
)

type taskController struct {
	taskInteractor interactor.TaskInteractor
}

func NewTaskController(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *taskController {
	taskRepository := gateway.NewTaskRepository(sqlhandler)
	taskInteractor := interactor.NewTastInteractor(sqlhandler, taskRepository, validator)

	return &taskController{taskInteractor}
}

func (con *taskController) Index(c interfaces.Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("id must be integer err: ", err)
		c.JSON(400, "bad request id must be integer", nil)
		return
	}

	tasks, err := con.taskInteractor.GetList(&interactor.TaskGetListInputDS{ProjectID: projectID})
	if err != nil {
		fmt.Println("usecase GetList error: ", err)
		c.JSON(500, "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, "ok", tasks)
}

func (con *taskController) Create(c interfaces.Context) {
	projectID, _ := strconv.Atoi(c.Param("id"))

	var data interactor.TaskStoreInputDS
	if err := c.Bind(&data); err != nil {
		c.JSON(400, "bad request", nil)
		return
	}

	data.ProjectID = projectID
	id, err := con.taskInteractor.Store(&data)
	if err != nil {
		fmt.Println("task usecase store error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			c.JSON(400, err.Error(), nil)
		default:
			c.JSON(500, err.Error(), nil)
		}
		return
	}

	c.JSON(200, "ok", map[string]interface{}{"id": id})
}

func (con *taskController) Update(c interfaces.Context) {
	taskID, _ := strconv.Atoi(c.Param("task_id"))

	var data interactor.TaskUpdateInputDS
	if err := c.Bind(&data); err != nil {
		fmt.Println("data decode error: ", err)
		c.JSON(400, err.Error(), nil)
		return
	}

	data.TaskID = taskID
	updatedTask, err := con.taskInteractor.Update(&data)
	if err != nil {
		fmt.Println("task usecase update error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			c.JSON(400, err.Error(), nil)
		default:
			c.JSON(500, err.Error(), nil)
		}
		return
	}

	c.JSON(200, "ok", map[string]interface{}{"task": updatedTask})
}

func (con *taskController) Delete(c interfaces.Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		fmt.Println("id must be integer err: ", err)
		c.JSON(400, "bad request id must be integer", nil)
		return
	}

	err = con.taskInteractor.Delete(&interactor.TaskDeleteInputDS{TaskID: taskID})
	if err != nil {
		fmt.Println("usecase Delete error: ", err)
		switch err.(type) {
		case *errors.NotFoundErr:
			c.JSON(404, err.Error(), nil)
		default:
			c.JSON(500, "Internal server error", nil)
		}
		return
	}

	c.JSON(204, "ok", nil)
}
