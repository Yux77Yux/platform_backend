package model

type UserLogin struct {
	UserDefault
	UserAvatar string `json:"user_avatar,omitempty"`
}
