package model

type Tag struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" validate:"required,max=30"`
	Color     string `json:"color" validate:"required,min=7,max=7"`
	ProjectID int    `json:"projectId"`
}

type Tags []Tag
