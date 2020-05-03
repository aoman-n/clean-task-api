package usecase

import (
	"task-api/src/entity/model"
	"task-api/src/utils/errors"
	"time"
)

type TaskInteractor interface {
	Store(*TaskStoreInputDS) (int64, error)
	Update(*TaskUpdateInputDS) (*model.Task, error)
	GetList(*TaskGetListInputDS) ([]*model.Task, error)
	Delete(*TaskDeleteInputDS) error
}

type taskInteractor struct {
	transactionHandler TransactionHandler
	taskRepository     TaskRepository
	validator          Validator
}

func NewTastInteractor(transactionHandler TransactionHandler, taskRepo TaskRepository, validator Validator) TaskInteractor {
	return &taskInteractor{transactionHandler, taskRepo, validator}
}

type TaskGetListInputDS struct {
	ProjectID int
}

func (ti *taskInteractor) GetList(in *TaskGetListInputDS) ([]*model.Task, error) {
	return ti.taskRepository.FetchByProjectID(nil, in.ProjectID)
}

type TaskStoreInputDS struct {
	ProjectID int
	Name      string `json:"name"`
}

func (ti *taskInteractor) Store(in *TaskStoreInputDS) (int64, error) {
	task := model.Task{
		Name:      in.Name,
		DueOn:     time.Now(),
		Status:    model.Waiting,
		ProjectID: in.ProjectID,
	}
	err := ti.validator.Struct(task)
	if err != nil {
		return 0, errors.NewModelValidationErr(err.Error())
	}

	id, err := ti.taskRepository.Create(nil, &task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

type TaskUpdateInputDS struct {
	TaskID int
	Name   *string `json:"name"`
	Status *int    `json:"status"`
}

func (ti *taskInteractor) Update(in *TaskUpdateInputDS) (*model.Task, error) {

	// updateするタスクを取得する
	task, err := ti.taskRepository.FindByID(nil, in.TaskID)
	if err != nil {
		return nil, err
	}

	// updateする項目を書き換える
	if in.Name != nil {
		task.Name = *in.Name
	}
	if in.Status != nil {
		task.Status = *in.Status
	}

	// validation
	err = ti.validator.Struct(*task)
	if err != nil {
		return nil, errors.NewModelValidationErr(err.Error())
	}

	// update
	_, err = ti.taskRepository.Save(nil, task)
	if err != nil {
		return nil, errors.NewRecordSaveErr(err.Error())
	}

	return task, nil
}

type TaskDeleteInputDS struct {
	TaskID int
}

func (ti *taskInteractor) Delete(in *TaskDeleteInputDS) error {
	_, err := ti.taskRepository.FindByID(nil, in.TaskID)
	if err != nil {
		return err
	}

	err = ti.taskRepository.Delete(nil, in.TaskID)
	if err != nil {
		return err
	}

	return nil
}
