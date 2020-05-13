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

type ProjectController struct {
	ProjectInteractor interactor.ProjectInteractor
}

func NewProjectController(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *ProjectController {
	userRepository := gateway.NewUserRepository(sqlhandler)
	projectRepository := gateway.NewProjectRepository(sqlhandler)
	projectInteractor := interactor.NewProjectInteractor(userRepository, projectRepository, sqlhandler, validator)

	return &ProjectController{
		ProjectInteractor: projectInteractor,
	}
}

func (uc *ProjectController) Index(c interfaces.Context) {
	uID := c.MustGet("userId").(int64)
	projects, err := uc.ProjectInteractor.GetList(&interactor.ProjectGetListInputDS{Uid: uID})
	if err != nil {
		fmt.Println("in project index, GetList error: ", err)
		c.JSON(500, err.Error(), nil)
		return
	}

	c.JSON(200, "ok", projects)
}

func (uc *ProjectController) Create(c interfaces.Context) {
	uID := c.MustGet("userId").(int64)
	var req interactor.ProjectStoreInputDS
	req.UserID = uID
	if err := c.Bind(&req); err != nil {
		fmt.Println("in project create, decode error: ", err)
		c.JSON(500, err.Error(), nil)
		return
	}

	fmt.Println("in project create. req: ", req)

	projectID, err := uc.ProjectInteractor.Store(&req)
	if err != nil {
		fmt.Println("in project create, store error: ", err)
		code, msg := uc.errStatus(err)
		c.JSON(code, msg, nil)
		return
	}

	c.JSON(201, "created", map[string]int64{"id": projectID})
}

func (uc *ProjectController) Delete(c interfaces.Context) {
	uID := c.MustGet("userId").(int64)
	projectID, _ := strconv.Atoi(c.Param("id"))

	err := uc.ProjectInteractor.Delete(&interactor.ProjectDeleteInputDS{
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
