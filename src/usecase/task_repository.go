package usecase

import "task-api/src/entity/model"

type TaskRepository interface {
	Create(tx Transaction, task *model.Task) (int64, error)
	FindByID(tx Transaction, id int) (*model.Task, error)
	Save(tx Transaction, task *model.Task) (int64, error)
}
