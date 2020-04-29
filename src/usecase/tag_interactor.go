package usecase

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/utils/errors"
)

type tagInteractor struct {
	transactionHandler TransactionHandler
	tagRepository      TagRepository
	projectRepository  ProjectRepository
	validator          Validator
}

type TagInteractor interface {
	Create(*TagCreateInputDS) (int64, error)
	GetList(*TagGetListInputDS) (*model.Tags, error)
	Delete(*TagDeleteInputDS) error
	Update(*TagUpdateInputDS) (*model.Tag, error)
}

func NewTagInteractor(
	txHandler TransactionHandler,
	tagRepo TagRepository,
	pRepo ProjectRepository,
	validator Validator,
) TagInteractor {
	return &tagInteractor{
		transactionHandler: txHandler,
		tagRepository:      tagRepo,
		projectRepository:  pRepo,
		validator:          validator,
	}
}

type TagCreateInputDS struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	ProjectID int
}

func (ti *tagInteractor) Create(in *TagCreateInputDS) (int64, error) {
	tag := model.Tag{
		Name:      in.Name,
		Color:     in.Color,
		ProjectID: in.ProjectID,
	}

	err := ti.validator.Struct(tag)
	if err != nil {
		return 0, errors.NewModelValidationErr(err.Error())
	}

	id, err := ti.tagRepository.Create(nil, &tag)
	if err != nil {
		return 0, errors.NewRecordSaveErr(err.Error())
	}

	return id, nil
}

type TagGetListInputDS struct {
	ProjectID int
}

func (ti *tagInteractor) GetList(in *TagGetListInputDS) (*model.Tags, error) {
	tags, err := ti.tagRepository.FetchByProjectID(nil, in.ProjectID)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

type TagUpdateInputDS struct {
	Name   *string `json:"name"`
	Color  *string `json:"color"`
	TagID  int
	UserID int64
}

func (ti *tagInteractor) Update(in *TagUpdateInputDS) (*model.Tag, error) {

	// tagの存在確認
	tag, err := ti.tagRepository.FindByID(nil, in.TagID)
	if err != nil {
		return nil, err
	}

	// role確認
	role, err := ti.projectRepository.RoleByProjectID(nil, in.UserID, tag.ProjectID)
	if err != nil {
		return nil, err
	}
	if role == model.Read {
		msg := fmt.Sprintf("not permission to project id %v", tag.ProjectID)
		return nil, errors.NewPermissionErr(msg)
	}

	// update値のセット
	if in.Name != nil {
		tag.Name = *in.Name
	}
	if in.Color != nil {
		tag.Color = *in.Color
	}
	// tagのvalidation
	err = ti.validator.Struct(tag)
	if err != nil {
		return nil, errors.NewModelValidationErr(err.Error())
	}

	// tagのUpdate
	err = ti.tagRepository.Save(nil, tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

type TagDeleteInputDS struct {
	TagID  int
	UserID int64
}

func (ti *tagInteractor) Delete(in *TagDeleteInputDS) error {

	// tagの存在確認
	tag, err := ti.tagRepository.FindByID(nil, in.TagID)
	if err != nil {
		return errors.NewNotFoundErr(err.Error())
	}

	// role確認
	role, err := ti.projectRepository.RoleByProjectID(nil, in.UserID, tag.ProjectID)
	if err != nil {
		return err
	}
	if role == model.Read {
		msg := fmt.Sprintf("not permission to project id %v", tag.ProjectID)
		return errors.NewPermissionErr(msg)
	}

	// 削除
	err = ti.tagRepository.Delete(nil, in.TagID)
	if err != nil {
		return err
	}

	return nil
}
