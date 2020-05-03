package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
)

type TagController struct {
	interactor usecase.TagInteractor
}

func NewTagController(sqlhandler SQLHandler, validator usecase.Validator) *TagController {
	tagRepository := NewTagRepository(sqlhandler)
	projectRepository := NewProjectRepository(sqlhandler)
	tagInteractor := usecase.NewTagInteractor(
		sqlhandler,
		tagRepository,
		projectRepository,
		validator,
	)

	return &TagController{tagInteractor}
}

func (con *TagController) Index(c Context) {
	projectID, _ := strconv.Atoi(c.Param("id"))
	var data usecase.TagCreateInputDS
	if err := c.Bind(&data); err != nil {
		c.JSON(400, "bad request", nil)
		return
	}

	tags, err := con.interactor.GetList(&usecase.TagGetListInputDS{ProjectID: projectID})
	if err != nil {
		fmt.Println("get tag list error: ", err)
		c.JSON(500, "Internal server error", nil)
		return
	}

	c.JSON(200, "ok", tags)
}

func (con *TagController) Create(c Context) {
	projectID, _ := strconv.Atoi(c.Param("id"))
	var data usecase.TagCreateInputDS
	if err := c.Bind(&data); err != nil {
		c.JSON(400, "bad request", nil)
		return
	}

	data.ProjectID = projectID
	id, err := con.interactor.Create(&data)
	if err != nil {
		fmt.Println("create tag error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			c.JSON(400, err.Error(), nil)
		default:
			c.JSON(500, "internal server error", nil)
		}
		return
	}

	c.JSON(200, "ok", map[string]interface{}{"id": id})
}

func (con *TagController) Update(c Context) {
	uID := c.MustGet("userId").(int64)
	tagID, _ := strconv.Atoi(c.Param("id"))
	var data usecase.TagUpdateInputDS
	if err := c.Bind(&data); err != nil {
		c.JSON(400, "bad request", nil)
		return
	}

	data.TagID = tagID
	data.UserID = uID
	tag, err := con.interactor.Update(&data)
	if err != nil {
		code, msg := con.errStatus(err)
		c.JSON(code, msg, nil)
		return
	}

	c.JSON(200, "ok", tag)
}

func (con *TagController) Delete(c Context) {
	uID := c.MustGet("userId").(int64)
	tagID, _ := strconv.Atoi(c.Param("id"))

	err := con.interactor.Delete(&usecase.TagDeleteInputDS{
		TagID:  tagID,
		UserID: uID,
	})
	if err != nil {
		fmt.Println("delete tag error: ", err)
		code, msg := con.errStatus(err)
		c.JSON(code, msg, nil)
		return
	}

	c.JSON(204, "deleted", nil)
}

func (con *TagController) errStatus(err error) (int, string) {
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
