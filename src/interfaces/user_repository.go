package interfaces

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/usecase"
)

type userRepository struct {
	sqlhandler SQLHandler
}

func NewUserRepository(sqlhandler SQLHandler) usecase.UserRepository {
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

func (repo *userRepository) FindByProjectID(projectID int) (*model.UserList, error) {
	query := `
	SELECT users.id,users.login_name,users.display_name,users.avatar_url,project_users.role
	FROM users
	INNER JOIN project_users
		ON users.id = project_users.user_id
	WHERE project_users.project_id = ?
	`
	rows, err := repo.sqlhandler.Query(query, projectID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var users model.UserList
	for rows.Next() {
		var u model.UserListItem
		if err := rows.Scan(&u.ID, &u.LoginName, &u.DisplayName, &u.AvatarURL, &u.Role); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (repo *userRepository) FindLikeLoginName(loginName string) (model.Users, error) {
	query := "SELECT * FROM users WHERE login_name LIKE '%" + loginName + "%'"

	rows, err := repo.sqlhandler.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users model.Users
	for rows.Next() {
		var u model.User
		err := rows.Scan(
			&u.ID,
			&u.DisplayName,
			&u.LoginName,
			&u.PasswordDigest,
			&u.AvatarURL,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
