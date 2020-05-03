package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
)

type taskController struct {
	taskInteractor usecase.TaskInteractor
}

func NewTaskController(sqlhandler SQLHandler, validator usecase.Validator) *taskController {
	taskRepository := NewTaskRepository(sqlhandler)
	taskInteractor := usecase.NewTastInteractor(sqlhandler, taskRepository, validator)

	return &taskController{taskInteractor}
}

func (con *taskController) Index(c Context) {
	projectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("id must be integer err: ", err)
		c.JSON(400, "bad request id must be integer", nil)
		return
	}

	tasks, err := con.taskInteractor.GetList(&usecase.TaskGetListInputDS{ProjectID: projectID})
	if err != nil {
		fmt.Println("usecase GetList error: ", err)
		c.JSON(500, "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, "ok", tasks)
}

func (con *taskController) Create(c Context) {
	projectID, _ := strconv.Atoi(c.Param("id"))

	var data usecase.TaskStoreInputDS
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

func (con *taskController) Update(c Context) {
	taskID, _ := strconv.Atoi(c.Param("task_id"))

	var data usecase.TaskUpdateInputDS
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

func (con *taskController) Delete(c Context) {
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		fmt.Println("id must be integer err: ", err)
		c.JSON(400, "bad request id must be integer", nil)
		return
	}

	err = con.taskInteractor.Delete(&usecase.TaskDeleteInputDS{TaskID: taskID})
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
