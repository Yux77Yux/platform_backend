package model

import "time"

type User struct {
	UserId     string     `json:"user_id"`
	UserName   string     `json:"user_name,omitempty"`
	UserAvatar string     `json:"user_avatar,omitempty"`
	UserGender UserGender `json:"user_gender,omitempty"`
	UserBio    string     `json:"user_bio,omitempty"`
	UserStatus UserStatus `json:"user_status"`
	UserBday   time.Time  `json:"user_bday,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
