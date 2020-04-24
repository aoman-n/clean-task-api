package usecase

import "task-api/src/entity/model"

type ProjectRepository interface {
	Create(Transaction, *model.Project) (int64, error)
	AddUser(userID int64, projectID int64, role string) (int64, error)
}
