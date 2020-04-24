package model

type Project struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" validate:"required,max=50"`
	Description string `json:"description" validate:"required,max=300"`
}
