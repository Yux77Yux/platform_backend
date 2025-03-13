package test

import (
	"github.com/Yux77Yux/platform_backend/generated/auth"
)

type Id struct {
	Id       string `json:"id"`
	Duration float64
}
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}
type Register_OK struct {
	*User
	Duration float64
}
type Register_ER struct {
	*User
	Error string `json:"error"`
}
type Login_ER struct {
	*User
	Error string `json:"error"`
}
type Login_OK struct {
	*User
	IdInDb       int64              `json:"idInDb"`
	RefreshToken *auth.RefreshToken `json:"token"`
	Duration     float64
}
type User_ER struct {
	*User
	Error string `json:"error"`
}

type Creation struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Bio        string `json:"bio"`
	Uid        string `json:"uid"`
	Src        string `json:"src"`
	Thumbnail  string `json:"thumbnail"`
	Duration   int32  `json:"duration"`
	CategoryId int32  `json:"categoryId"`
}
type Upload_OK struct {
	*Creation
	Duration float64
}
type Creation_ER struct {
	*Creation
	Error string `json:"error"`
}
type User_After_Upload struct {
	*User
	IdInDb int64 `json:"idInDb"`
}
