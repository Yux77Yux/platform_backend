package model

import "time"

type UserGender uint8
type UserStatus uint8

const (
	Unspecified UserGender = 0
	Male        UserGender = 1
	Female      UserGender = 2

	Hiding   UserStatus = 0 // 二进制: 00
	Inactive UserStatus = 1 // 二进制: 01
	Active   UserStatus = 2 // 二进制: 10
	Limited  UserStatus = 3 // 二进制: 11
)

type UserDefault struct {
	UserUuid   string `json:"user_uuid"`
	UserName   string `json:"user_name,omitempty"`
	UserAvatar string `json:"user_avatar,omitempty"`
}

type UserCreation struct {
	UserDefault
	UserBio string `json:"user_bio"`
}

type UserTalk struct {
	UserDefault
	UserStatus UserStatus `json:"user_status"`
}

type User struct {
	UserDefault
	UserGender UserGender `json:"user_gender,omitempty"`
	UserBio    string     `json:"user_bio,omitempty"`
	UserStatus UserStatus `json:"user_status"`
	UserBday   time.Time  `json:"user_bday,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
