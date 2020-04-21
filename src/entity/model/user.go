package model

import "time"

type User struct {
	ID             int       `json:"id"`
	DisplayName    string    `json:"displayName"`
	LoginName      string    `json:"loginName"`
	PasswordDigest string    `json:"-"`
	AvatarURL      string    `json:"avatarUrl"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
