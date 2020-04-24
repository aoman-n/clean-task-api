package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/src/usecase"
)

type ProjectController struct {
	ProjectInteractor usecase.ProjectInteractor
}

func NewProjectController(sqlhandler SQLHandler, validator usecase.Validator) *ProjectController {
	userRepository := NewUserRepository(sqlhandler)
	projectRepository := NewProjectRepository(sqlhandler)
	projectInteractor := usecase.NewProjectInteractor(userRepository, projectRepository, sqlhandler)

	return &ProjectController{
		ProjectInteractor: projectInteractor,
	}
}

func (uc *ProjectController) Create(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {

	var data usecase.ProjectStoreInputDS
	data.UserID = uID
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("in project create, decode error: ", err)
		jsonView(w, 400, "bad request")
		return
	}

	fmt.Println("in project create. data: ", data)

	projectID, err := uc.ProjectInteractor.Store(&data)
	if err != nil {
		fmt.Println("in project create, store error: ", err)
		jsonView(w, 500, "server error")
		return
	}

	jsonView(w, 200, fmt.Sprintf("hello, project, id is %v", projectID))
}
