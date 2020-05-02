package interfaces

import (
	"fmt"
	"task-api/src/entity/model"
	"task-api/src/usecase"
	"task-api/src/utils/errors"
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

func (repo *tagRepository) FetchByProjectID(tx usecase.Transaction, projectID int) ([]*model.Tag, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	query := `SELECT * FROM tags WHERE project_id=?`
	rows, err := sqlhandler.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]*model.Tag, 0)
	for rows.Next() {
		t := new(model.Tag)
		err := rows.Scan(&t.ID, &t.Name, &t.Color, &t.ProjectID)
		if err != nil {
			return nil, err
		}

		tags = append(tags, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *tagRepository) Save(tx usecase.Transaction, t *model.Tag) error {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	query := `UPDATE tags SET name=?,color=? WHERE id=?`
	result, err := sqlhandler.Exec(query, t.Name, t.Color, t.ID)
	if err != nil {
		return errors.NewRecordSaveErr(err.Error())
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return errors.NewRecordSaveErr(err.Error())
	}
	if affect != 1 {
		return errors.NewRecordSaveErr(err.Error())
	}

	return nil
}

func (repo *tagRepository) FindByID(tx usecase.Transaction, id int) (*model.Tag, error) {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	var tag model.Tag
	err := sqlhandler.
		QueryRow(`SELECT * FROM tags WHERE id=?`, id).
		Scan(&tag.ID, &tag.Name, &tag.Color, &tag.ProjectID)

	if err != nil {
		return nil, errors.NewNotFoundErr(fmt.Sprintf("not found tag. id=%d", id))
	}

	return &tag, nil
}

func (repo *tagRepository) Delete(tx usecase.Transaction, id int) error {
	sqlhandler := repo.sqlhandler.FromTransaction(tx)

	_, err := sqlhandler.Exec(`DELETE FROM tags WHERE id=?`, id)
	if err != nil {
		return err
	}

	return nil
}
