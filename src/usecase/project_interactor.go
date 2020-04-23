package usecase

import (
	"fmt"
	"task-api/src/entity/model"
)

type ProjectInteractor interface {
	Store(*ProjectStoreInputDS) (int64, error)
}

type projectInteractor struct {
	UserRepository    UserRepository
	ProjectRepository ProjectRepository
}

func NewProjectInteractor(userRepo UserRepository, projectRepo ProjectRepository) ProjectInteractor {
	return &projectInteractor{userRepo, projectRepo}
}

type ProjectStoreInputDS struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int64
}

func (pi *projectInteractor) Store(in *ProjectStoreInputDS) (int64, error) {

	// TODO: バリデーション
	project := &model.Project{Title: in.Title, Description: in.Description}
	// TODO: トランザクション開始
	projectID, err := pi.ProjectRepository.Create(project)
	if err != nil {
		fmt.Println("Create error: ", err)
		// TODO: rollback
		return 0, nil
	}

	_, err = pi.ProjectRepository.AddUser(in.UserID, projectID, model.Admin)
	if err != nil {
		fmt.Println("AddUser error: ", err)
		// TODO: rollback
		return 0, nil
	}

	return projectID, nil
}
