package interactor

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/entity/repository"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
)

type ProjectInteractor interface {
	Store(*ProjectStoreInputDS) (int64, error)
	GetList(*ProjectGetListInputDS) ([]*model.ProjectResult, error)
	Delete(*ProjectDeleteInputDS) error
}

type projectInteractor struct {
	UserRepository     repository.UserRepository
	ProjectRepository  repository.ProjectRepository
	TransactionHandler repository.TransactionHandler
	Validator          usecase.Validator
}

func NewProjectInteractor(
	userRepo repository.UserRepository,
	projectRepo repository.ProjectRepository,
	transactionHandler repository.TransactionHandler,
	validator usecase.Validator,
) ProjectInteractor {
	return &projectInteractor{userRepo, projectRepo, transactionHandler, validator}
}

type ProjectStoreInputDS struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64
}

func (pi *projectInteractor) Store(in *ProjectStoreInputDS) (int64, error) {
	project := &model.Project{Title: in.Title, Description: in.Description}
	err := pi.Validator.Struct(project)
	if err != nil {
		return 0, errors.NewModelValidationErr(err.Error())
	}

	// トランザクション内で、project作成とアドミンUserとして参加させる処理
	projectID, err := pi.TransactionHandler.TransactAndReturnData(func(tx repository.Transaction) (interface{}, error) {
		projectID, err := pi.ProjectRepository.Create(tx, project)
		if err != nil {
			return nil, err
		}

		return pi.ProjectRepository.AddUser(tx, in.UserID, projectID, model.Admin)
	})

	if err != nil {
		fmt.Println("Project Store error: ", err)
		return 0, err
	}

	return projectID.(int64), nil
}

type ProjectGetListInputDS struct {
	Uid int64
}

func (pi *projectInteractor) GetList(in *ProjectGetListInputDS) ([]*model.ProjectResult, error) {
	projects, err := pi.ProjectRepository.FindByUserID(nil, in.Uid)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

type ProjectDeleteInputDS struct {
	Uid       int64
	ProjectID int
}

func (pi *projectInteractor) Delete(in *ProjectDeleteInputDS) error {
	role, err := pi.ProjectRepository.RoleByProjectID(nil, in.Uid, in.ProjectID)
	if err != nil {
		return err
	}

	if role != model.Admin {
		return errors.NewPermissionErr(fmt.Sprintf("%s cannot delete permission", role))
	}

	return pi.ProjectRepository.Delete(nil, in.ProjectID)
}
