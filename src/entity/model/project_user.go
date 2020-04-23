package model

type ProjectUser struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	ProjectID int64  `json:"project_id"`
	Role      string `json:"role"` // admin, write, read
}

// Role constants
const (
	Admin = "admin"
	Write = "write"
	Read  = "read"
)
