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

func (repo *taskRepository) FetchByProjectID(tx usecase.Transaction, pID int) (*model.Tasks, error) {
	sqlhandler := repo.FromTransaction(tx)

	rows, err := sqlhandler.Query(`SELECT * FROM tasks WHERE project_id=?`, pID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks model.Tasks
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.Name, &t.DueOn, &t.Status, &t.ProjectID)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (repo *taskRepository) Delete(tx usecase.Transaction, tID int) error {
	sqlhandler := repo.FromTransaction(tx)

	result, err := sqlhandler.Exec(`DELETE FROM tasks WHERE id=?`, tID)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
	}

	return nil
}
