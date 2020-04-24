package usecase

import (
	"fmt"
	"task-api/src/entity/model"
)

type ProjectInteractor interface {
	Store(*ProjectStoreInputDS) (int64, error)
}

type projectInteractor struct {
	UserRepository     UserRepository
	ProjectRepository  ProjectRepository
	TransactionHandler TransactionHandler
}

func NewProjectInteractor(userRepo UserRepository, projectRepo ProjectRepository, transactionHandler TransactionHandler) ProjectInteractor {
	return &projectInteractor{userRepo, projectRepo, transactionHandler}
}

type ProjectStoreInputDS struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64
}

func (pi *projectInteractor) Store(in *ProjectStoreInputDS) (int64, error) {
	// TODO: バリデーション
	project := &model.Project{Title: in.Title, Description: in.Description}

	// トランザクション内で、project作成とアドミンUserとして参加させる処理
	projectID, err := pi.TransactionHandler.TransactAndReturnData(func(tx Transaction) (interface{}, error) {
		projectID, err := pi.ProjectRepository.Create(tx, project)
		if err != nil {
			return nil, err
		}

		return pi.ProjectRepository.AddUser(in.UserID, projectID, model.Admin)
	})

	if err != nil {
		fmt.Println("Project Store error: ", err)
		return 0, err
	}

	return projectID.(int64), nil
}
