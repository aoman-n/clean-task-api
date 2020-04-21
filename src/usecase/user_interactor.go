package usecase

import "task-api/src/entity/model"

type userInteractor struct {
	userRepository UserRepository
	validator      Validator
}

type UserInteractor interface {
	Create(*model.User) (int, error)
}

func NewUserInteractor(repo UserRepository, validator Validator) UserInteractor {
	return &userInteractor{repo, validator}
}

func (ui *userInteractor) Create(user *model.User) (id int, err error) {
	err = ui.validator.Struct(user)
	if err != nil {
		return
	}
	id, err = ui.userRepository.Store(user)
	return
}

func (ui *userInteractor) User(id int) (user *model.User, err error) {
	user, err = ui.userRepository.FindByID(id)
	return
}
