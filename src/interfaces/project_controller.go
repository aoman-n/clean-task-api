package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
)

type ProjectController struct {
	ProjectInteractor usecase.ProjectInteractor
}

func NewProjectController(sqlhandler SQLHandler, validator usecase.Validator) *ProjectController {
	userRepository := NewUserRepository(sqlhandler)
	projectRepository := NewProjectRepository(sqlhandler)
	projectInteractor := usecase.NewProjectInteractor(userRepository, projectRepository, sqlhandler, validator)

	return &ProjectController{
		ProjectInteractor: projectInteractor,
	}
}

func (uc *ProjectController) Index(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	projects, err := uc.ProjectInteractor.GetList(&usecase.ProjectGetListInputDS{Uid: uID})
	if err != nil {
		fmt.Println("in project index, GetList error: ", err)
		jsonView(w, 500, err.Error())
		return
	}

	jsonView(w, 200, projects)
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
		switch err.(type) {
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
			return
		default:
			jsonView(w, 500, err.Error())
			return
		}
	}

	jsonView(w, 200, fmt.Sprintf("hello, project, id is %v", projectID))
}
