package usecase

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/utils/errors"
)

type ProjectInteractor interface {
	Store(*ProjectStoreInputDS) (int64, error)
}

type projectInteractor struct {
	UserRepository     UserRepository
	ProjectRepository  ProjectRepository
	TransactionHandler TransactionHandler
	Validator          Validator
}

func NewProjectInteractor(userRepo UserRepository, projectRepo ProjectRepository, transactionHandler TransactionHandler, validator Validator) ProjectInteractor {
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
	projectID, err := pi.TransactionHandler.TransactAndReturnData(func(tx Transaction) (interface{}, error) {
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
