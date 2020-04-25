package interfaces

import (
	"task-api/src/entity/model"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
)

type taskRepository struct {
	SQLHandler
}

func NewTaskRepository(sqlhandler SQLHandler) usecase.TaskRepository {
	return &taskRepository{sqlhandler}
}

func (repo *taskRepository) Create(tx usecase.Transaction, t *model.Task) (int64, error) {
	sqlhandler := repo.FromTransaction(tx)

	query := `INSERT INTO tasks (name, due_on, status, project_id) VALUES (?, ?, ?, ?)`
	result, err := sqlhandler.Exec(query, t.Name, t.DueOn, t.Status, t.ProjectID)
	if err != nil {
		return 0, errors.NewRecordSaveErr("failed to create task")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.NewRecordSaveErr("failed to create task")
	}

	return id, nil
}
