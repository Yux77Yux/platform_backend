package model

type UserTalk struct {
	UserDefault
	UserStatus UserStatus `json:"user_status"`
}
