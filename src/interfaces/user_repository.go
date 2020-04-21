package interfaces

import (
	"fmt"
	"task-api/src/entity/model"
)

type userRepository struct {
	sqlhandler SQLHandler
}

func NewUserRepository(sqlhandler SQLHandler) *userRepository {
	return &userRepository{sqlhandler}
}

func (repo *userRepository) Store(u *model.User) (id int64, err error) {
	query := `
	INSERT INTO users
		(display_name, login_name, password_digest)
	VALUES (?, ?, ?)
	`
	result, err := repo.sqlhandler.Exec(query, u.LoginName, u.LoginName, u.PasswordDigest)
	if err != nil {
		fmt.Println("insert err ", err)
		return 0, err
	}
	id, err = result.LastInsertId()
	return
}

func (repo *userRepository) FindByID(id int) (user *model.User, err error) {
	return
}
