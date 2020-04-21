package usecase

import (
	"fmt"
	"task-api/src/entity/model"

	"golang.org/x/crypto/bcrypt"
)

type userInteractor struct {
	userRepository UserRepository
	validator      Validator
}

type UserInteractor interface {
	Store(UserStoreInputDS) (int64, error)
}

type UserStoreInputDS struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

func NewUserInteractor(repo UserRepository, validator Validator) UserInteractor {
	return &userInteractor{repo, validator}
}

func (ui *userInteractor) Store(input UserStoreInputDS) (id int64, err error) {

	passDigest, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := model.User{
		LoginName:      input.LoginName,
		PasswordDigest: string(passDigest),
	}
	err = ui.validator.Struct(user)
	if err != nil {
		fmt.Print("validation error ", err)
		return
	}
	id, err = ui.userRepository.Store(&user)
	if err != nil {
		fmt.Print("store error ", err)
	}
	return
}

func (ui *userInteractor) User(id int) (user *model.User, err error) {
	user, err = ui.userRepository.FindByID(id)
	return
}

// bcrypt.CompareHashAndPassword(passB, []byte(data.Password))
