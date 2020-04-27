package interfaces

import (
	"fmt"
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

func (repo *taskRepository) FindByID(tx usecase.Transaction, id int) (*model.Task, error) {
	sqlhandler := repo.FromTransaction(tx)

	var task model.Task
	err := sqlhandler.
		QueryRow(`SELECT * FROM tasks WHERE id = ?`, id).
		Scan(&task.ID, &task.Name, &task.DueOn, &task.Status, &task.ProjectID)

	if err != nil {
		fmt.Println("FindByID err: ", err)
		return nil, err
	}

	return &task, nil
}

func (repo *taskRepository) Save(tx usecase.Transaction, t *model.Task) (int64, error) {
	sqlhandler := repo.FromTransaction(tx)

	query := `UPDATE tasks SET name=?, due_on=?, status=? WHERE id=?`
	result, err := sqlhandler.Exec(query, t.Name, t.DueOn, t.Status, t.ID)
	if err != nil {
		return 0, errors.NewRecordSaveErr(err.Error())
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return 0, errors.NewRecordSaveErr(err.Error())
	}
	if affect != 1 {
		return 0, fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
	}

	return t.ID, nil
}
