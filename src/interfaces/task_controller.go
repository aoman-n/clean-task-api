package interfaces

import (
	"encoding/json"
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

func (con *taskController) Create(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	projectID, _ := strconv.Atoi(ps.ByName("id"))

	var data usecase.TaskStoreInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		jsonView(w, 400, "bad request")
		return
	}

	data.ProjectID = projectID
	id, err := con.taskInteractor.Store(&data)
	if err != nil {
		fmt.Println("task usecase store error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
		default:
			jsonView(w, 500, err.Error())
		}
		return
	}

	jsonView(w, 200, map[string]interface{}{"id": id})
}

func (con *taskController) Update(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	taskID, _ := strconv.Atoi(ps.ByName("task_id"))

	var data usecase.TaskUpdateInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("data decode error: ", err)
		jsonView(w, 400, "bad request")
		return
	}
	fmt.Printf("data, Name: %v, Status: %v", *data.Name, *data.Status)
	data.TaskID = taskID

	updatedTask, err := con.taskInteractor.Update(&data)
	if err != nil {
		fmt.Println("task usecase update error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
		default:
			jsonView(w, 500, err.Error())
		}
		return
	}

	jsonView(w, 200, map[string]interface{}{"task": updatedTask})
}
