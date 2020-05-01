package interfaces

import (
	"fmt"
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
		switch err.(type) {
		case *errors.ModelValidationErr:
			c.JSON(400, err.Error(), nil)
			return
		default:
			c.JSON(500, err.Error(), nil)
			return
		}
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
		switch err.(type) {
		case *errors.NotFoundErr:
			c.JSON(404, err.Error(), nil)
		case *errors.PermissionErr:
			c.JSON(401, err.Error(), nil)
		default:
			c.JSON(500, "Internal server error", nil)
		}
		return
	}

	c.JSON(204, "deleted", nil)
}
