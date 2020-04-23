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
	FindByLoginNameAndVerifyPassword(input UserLoginInputDS) (int64, error)
}

func NewUserInteractor(repo UserRepository, validator Validator) UserInteractor {
	return &userInteractor{repo, validator}
}

type UserStoreInputDS struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
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

type UserLoginInputDS struct {
	LoingName string `json:"loginName"`
	Password  string `json:"password"`
}

func (ui *userInteractor) FindByLoginNameAndVerifyPassword(input UserLoginInputDS) (int64, error) {
	user, err := ui.userRepository.FindByLoginName(input.LoingName)
	if err != nil {
		return 0, fmt.Errorf("not found user. loginName is %s", input.LoingName)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(input.Password))
	if err != nil {
		return 0, fmt.Errorf("password is invalid")
	}

	return user.ID, nil
}
