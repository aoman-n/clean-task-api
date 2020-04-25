package usecase

import "task-api/src/entity/model"

type TaskRepository interface {
	Create(tx Transaction, task *model.Task) (id int64, err error)
}
