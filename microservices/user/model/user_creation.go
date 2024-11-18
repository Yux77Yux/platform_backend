package model

type UserCreation struct {
	UserDefault
	UserBio string `json:"user_bio"`
}
