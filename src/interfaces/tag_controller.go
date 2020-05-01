package interfaces

import (
	"fmt"
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
			c.JSON(500, "Internal server error", nil)
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
		fmt.Println("update tag error: ", err)
		var code int
		switch err.(type) {
		case *errors.NotFoundErr:
			code = 404
		case *errors.ModelValidationErr:
			code = 400
		case *errors.PermissionErr:
			code = 401
		default:
			code = 500
		}
		c.JSON(code, err.Error(), nil)
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
		var code int
		switch err.(type) {
		case *errors.NotFoundErr:
			code = 404
		case *errors.ModelValidationErr:
			code = 400
		case *errors.PermissionErr:
			code = 401
		default:
			code = 500
		}
		c.JSON(code, err.Error(), nil)
		return
	}

	c.JSON(204, "deleted", nil)
}
