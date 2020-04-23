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

func (repo *projectRepository) Create(p *model.Project) (int64, error) {
	qeury := `insert into projects (title, description) values (?, ?)`
	result, err := repo.sqlhandler.Exec(qeury, p.Title, p.Description)
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

func (repo *projectRepository) AddUser(userID int64, projectID int64, role string) (int64, error) {
	qeury := `insert into project_users (user_id, project_id, role) values (?, ?, ?)`
	result, err := repo.sqlhandler.Exec(qeury, userID, projectID, role)
	if err != nil {
		fmt.Println("projcet_users create error: ", err)
		return 0, err
	}

	projectID, err = result.LastInsertId()
	if err != nil {
		fmt.Println("project_users create LastInsertId error: ", err)
	}

	return projectID, nil
}
