package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
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

func (uc *ProjectController) Index(c Context) {
	uID := c.MustGet("userId").(int64)
	projects, err := uc.ProjectInteractor.GetList(&usecase.ProjectGetListInputDS{Uid: uID})
	if err != nil {
		fmt.Println("in project index, GetList error: ", err)
		c.JSON(500, err.Error(), nil)
		return
	}

	c.JSON(200, "ok", projects)
}

func (uc *ProjectController) Create(c Context) {
	uID := c.MustGet("userId").(int64)
	var data usecase.ProjectStoreInputDS
	data.UserID = uID
	if err := c.Bind(&data); err != nil {
		fmt.Println("in project create, decode error: ", err)
		c.JSON(500, err.Error(), nil)
		return
	}

	fmt.Println("in project create. data: ", data)

	projectID, err := uc.ProjectInteractor.Store(&data)
	if err != nil {
		fmt.Println("in project create, store error: ", err)
		code, msg := uc.errStatus(err)
		c.JSON(code, msg, nil)
		return
	}

	c.JSON(201, "created", map[string]int64{"id": projectID})
}

func (uc *ProjectController) Delete(c Context) {
	uID := c.MustGet("userId").(int64)
	projectID, _ := strconv.Atoi(c.Param("id"))

	err := uc.ProjectInteractor.Delete(&usecase.ProjectDeleteInputDS{
		Uid:       uID,
		ProjectID: projectID,
	})

	if err != nil {
		fmt.Println("in project delete, delete error: ", err)
		code, msg := uc.errStatus(err)
		c.JSON(code, msg, nil)
		return
	}

	c.JSON(204, "deleted", nil)
}

func (uc *ProjectController) errStatus(err error) (int, string) {
	switch err.(type) {
	case *errors.PermissionErr:
		return http.StatusUnauthorized, err.Error()
	case *errors.NotFoundErr:
		return http.StatusNotFound, err.Error()
	case *errors.ModelValidationErr:
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
