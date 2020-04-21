package model

type ProjectUser struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProjectID int    `json:"project_id"`
	Role      string `json:"role"` // admin, write, read
}
