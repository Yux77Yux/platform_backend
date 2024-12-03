package model

type UserDefault struct {
	UserId     string `json:"user_id"`
	UserName   string `json:"user_name,omitempty"`
	UserAvatar string `json:"user_avatar,omitempty"`
}
