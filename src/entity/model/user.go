package model

import "time"

type User struct {
	ID             int64     `json:"id"`
	DisplayName    string    `json:"displayName"`
	LoginName      string    `json:"loginName" validate:"required"`
	PasswordDigest string    `json:"-"`
	AvatarURL      string    `json:"avatarUrl"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type Users []User

type UserListItem struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"displayName"`
	LoginName   string `json:"loginName" validate:"required"`
	AvatarURL   string `json:"avatarUrl"`
	Role        string `json:"role"`
}

type UserList []UserListItem
