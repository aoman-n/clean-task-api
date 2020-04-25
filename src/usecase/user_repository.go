package usecase

import "task-api/src/entity/model"

type UserRepository interface {
	Store(user *model.User) (id int64, err error)
	FindByID(id int) (user *model.User, err error)
	FindByLoginName(loginName string) (user *model.User, err error)
	FindByProjectID(userID int) (*model.Users, error)
}
