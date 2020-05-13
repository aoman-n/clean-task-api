package gateway

import (
	"fmt"
	"strings"
	"task-api/src/entity/model"
	"task-api/src/entity/repository"
	"task-api/src/interfaces"
	"task-api/src/utils/errors"
)

type taskRepository struct {
	interfaces.SQLHandler
}

func NewTaskRepository(sqlhandler interfaces.SQLHandler) repository.TaskRepository {
	return &taskRepository{sqlhandler}
}

func (repo *taskRepository) Create(tx repository.Transaction, t *model.Task) (int64, error) {
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
	t.ID = id

	return id, nil
}

func (repo *taskRepository) FindByID(tx repository.Transaction, id int) (*model.Task, error) {
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

func (repo *taskRepository) Save(tx repository.Transaction, t *model.Task) (int64, error) {
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

func (repo *taskRepository) FetchByProjectID(tx repository.Transaction, pID int) ([]*model.Task, error) {
	sqlhandler := repo.FromTransaction(tx)

	rows, err := sqlhandler.Query(`SELECT * FROM tasks WHERE project_id=?`, pID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*model.Task, 0)
	for rows.Next() {
		t := new(model.Task)
		err := rows.Scan(&t.ID, &t.Name, &t.DueOn, &t.Status, &t.ProjectID)
		if err != nil {
			return nil, err
		}
		t.Tags = []*model.Tag{}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	taskIDs := make([]int64, 0)
	for _, t := range tasks {
		taskIDs = append(taskIDs, t.ID)
	}

	tags, err := repo.fetchTags(nil, taskIDs)
	if err != nil {
		return nil, err
	}

	// TODO: もっと効率よく
	for _, tag := range tags {
		for _, task := range tasks {
			if tag.TaskID == task.ID {
				task.Tags = append(task.Tags, tag)
			}
		}
	}

	return tasks, nil
}

func (repo *taskRepository) fetchTags(tx repository.Transaction, taskIDs []int64) ([]*model.Tag, error) {
	sqlhandler := repo.FromTransaction(tx)

	query := `
	SELECT t.id,t.name,t.color,tt.task_id FROM task_tags as tt
		INNER JOIN tags as t ON tt.tag_id = t.id
		WHERE tt.task_id IN (?` + strings.Repeat(",?", len(taskIDs)-1) + `)`

	// interface{}型にしないと展開して引数に渡せないので変換
	ids := make([]interface{}, len(taskIDs))
	for i, v := range taskIDs {
		ids[i] = v
	}

	rows, err := sqlhandler.Query(query, ids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*model.Tag
	for rows.Next() {
		var t model.Tag
		rows.Scan(&t.ID, &t.Name, &t.Color, &t.TaskID)

		if err != nil {
			return nil, err
		}

		tags = append(tags, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *taskRepository) Delete(tx repository.Transaction, tID int) error {
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
