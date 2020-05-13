package model

import "time"

type Task struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,max=255"`
	DueOn     time.Time `json:"-"`
	Status    int       `json:"status"`
	ProjectID int       `json:"-"`
}

const (
	Waiting int = iota + 1
	Doing
	Done
	Canceled
)
