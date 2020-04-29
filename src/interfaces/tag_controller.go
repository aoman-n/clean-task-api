package interfaces

import (
	"encoding/json"
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

func (con *TagController) Index(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	projectID, _ := strconv.Atoi(ps.ByName("id"))
	var data usecase.TagCreateInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		jsonView(w, 400, "bad request")
		return
	}

	tags, err := con.interactor.GetList(&usecase.TagGetListInputDS{ProjectID: projectID})
	if err != nil {
		fmt.Println("get tag list error: ", err)
		jsonView(w, 500, "Internal server error")
		return
	}

	jsonView(w, 200, map[string]interface{}{"tags": tags})
}

func (con *TagController) Create(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	projectID, _ := strconv.Atoi(ps.ByName("id"))
	var data usecase.TagCreateInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		jsonView(w, 400, "bad request")
		return
	}

	data.ProjectID = projectID
	id, err := con.interactor.Create(&data)
	if err != nil {
		fmt.Println("create tag error: ", err)
		switch err.(type) {
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
		default:
			jsonView(w, 500, "Internal server error")
		}
		return
	}
	jsonView(w, 200, map[string]interface{}{"id": id})
}

func (con *TagController) Update(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	tagID, _ := strconv.Atoi(ps.ByName("id"))
	var data usecase.TagUpdateInputDS
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		jsonView(w, 400, "bad request")
		return
	}

	data.TagID = tagID
	data.UserID = uID
	tag, err := con.interactor.Update(&data)
	if err != nil {
		fmt.Println("update tag error: ", err)
		switch err.(type) {
		case *errors.NotFoundErr:
			jsonView(w, 404, err.Error())
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
		case *errors.PermissionErr:
			jsonView(w, 401, err.Error())
		default:
			jsonView(w, 500, err.Error())
		}
		return
	}

	jsonView(w, 200, tag)
}

func (con *TagController) Delete(w http.ResponseWriter, r *http.Request, ps Params, uID int64) {
	tagID, _ := strconv.Atoi(ps.ByName("id"))

	err := con.interactor.Delete(&usecase.TagDeleteInputDS{
		TagID:  tagID,
		UserID: uID,
	})
	if err != nil {
		fmt.Println("delete tag error: ", err)
		switch err.(type) {
		case *errors.NotFoundErr:
			jsonView(w, 404, err.Error())
		case *errors.ModelValidationErr:
			jsonView(w, 400, err.Error())
		case *errors.PermissionErr:
			jsonView(w, 401, err.Error())
		default:
			jsonView(w, 500, err.Error())
		}
		return
	}

	jsonView(w, 204, "deleted")
}
