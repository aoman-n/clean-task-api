package usecase

import "task-api/src/entity/model"

type TagRepository interface {
	Create(tx Transaction, tag *model.Tag) (int64, error)
	FetchByProjectID(tx Transaction, projectID int) (*model.Tags, error)
	Save(tx Transaction, tag *model.Tag) (*model.Tag, error)
}
