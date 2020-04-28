package usecase

import (
	"task-api/src/entity/model"
	"task-api/src/utils/errors"
)

type tagInteractor struct {
	transactionHandler TransactionHandler
	tagRepository      TagRepository
	validator          Validator
}

type TagInteractor interface {
	Create(*TagCreateInputDS) (int64, error)
	GetList(*TagGetListInputDS) (*model.Tags, error)
}

func NewTagInteractor(txHandler TransactionHandler, tagRepo TagRepository, validator Validator) TagInteractor {
	return &tagInteractor{
		transactionHandler: txHandler,
		tagRepository:      tagRepo,
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
	ProjectID int
	UserID    int
}

func (ti *tagInteractor) Update(in *TagUpdateInputDS) (*model.Tag, error) {

	// tagの存在確認
	// role確認
	// tagのvalidation
	// tagのUpdate

	return nil, nil
}
