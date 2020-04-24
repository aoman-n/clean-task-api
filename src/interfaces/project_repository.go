package interfaces

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/usecase"
)

type projectRepository struct {
	sqlhandler SQLHandler
}

func NewProjectRepository(sqlhandler SQLHandler) usecase.ProjectRepository {
	return &projectRepository{sqlhandler}
}

func (repo *projectRepository) Create(tx usecase.Transaction, p *model.Project) (int64, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	qeury := `insert into projects (title, description) values (?, ?)`
	result, err := sqlhandler.Exec(qeury, p.Title, p.Description)
	if err != nil {
		fmt.Println("projcet creat error: ", err)
		return 0, err
	}

	projectID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("project create LastInsertId error: ", err)
	}

	return projectID, nil
}

func (repo *projectRepository) AddUser(tx usecase.Transaction, userID int64, projectID int64, role string) (int64, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	qeury := `insert into project_users (user_id, project_id, role) values (?, ?, ?)`
	_, err := sqlhandler.Exec(qeury, userID, projectID, role)
	if err != nil {
		fmt.Println("projcet_users create error: ", err)
		return 0, err
	}

	return projectID, nil
}

func (repo *projectRepository) FindByUserID(tx usecase.Transaction, userID int64) (model.ProjectResults, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	query := `
	SELECT
		projects.id,
		projects.title,
		projects.description,
		project_users.role
	FROM
		projects
	JOIN
		project_users ON projects.id = project_users.project_id
	WHERE
		project_users.user_id = ?
	`
	rows, err := sqlhandler.Query(query, userID)
	defer rows.Close()

	if err != nil {
		fmt.Println("FindByUserID error: ", err)
		return nil, err
	}

	var projects model.ProjectResults
	for rows.Next() {
		var p model.ProjectResult
		if err = rows.Scan(&p.ID, &p.Title, &p.Description, &p.Role); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
