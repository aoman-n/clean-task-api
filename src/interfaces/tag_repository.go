package interfaces

import (
	"task-api/src/entity/model"
	"task-api/src/usecase"
)

type tagRepository struct {
	sqlhandler SQLHandler
}

func NewTagRepository(sqlhandler SQLHandler) usecase.TagRepository {
	return &tagRepository{sqlhandler}
}

func (repo *tagRepository) Create(tx usecase.Transaction, t *model.Tag) (int64, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	query := `INSERT INTO tags (name,color,project_id) VALUES (?,?,?)`
	result, err := sqlhandler.Exec(query, t.Name, t.Color, t.ProjectID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *tagRepository) FetchByProjectID(tx usecase.Transaction, projectID int) (*model.Tags, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	query := `SELECT * FROM tags WHERE project_id=?`
	rows, err := sqlhandler.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags model.Tags
	for rows.Next() {
		var t model.Tag
		err := rows.Scan(&t.ID, &t.Name, &t.Color, &t.ProjectID)
		if err != nil {
			return nil, err
		}

		tags = append(tags, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tags, nil
}

func (repo *tagRepository) Save(tx usecase.Transaction, t *model.Tag) (*model.Tag, error) {
	return nil, nil
}
