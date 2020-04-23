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

func (repo *userRepository) FindByLoginName(loginName string) (*model.User, error) {
	query := `select * from users where login_name = ?`

	var user model.User
	err := repo.sqlhandler.QueryRow(query, loginName).Scan(
		&user.ID,
		&user.DisplayName,
		&user.LoginName,
		&user.PasswordDigest,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		fmt.Println("in FindByLoginNmame err: ", err)
		return nil, err
	}

	fmt.Println("in FindByLoginNmame user: ", user)
	return &user, nil
}
