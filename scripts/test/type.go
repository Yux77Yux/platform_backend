package test

import (
	"github.com/Yux77Yux/platform_backend/generated/auth"
	"github.com/Yux77Yux/platform_backend/scripts/data"
)

type Id struct {
	Id       string `json:"id"`
	Duration float64
}
type User = data.User
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

type Creation = data.Creation
type Creation_OK struct {
	*Creation
	IdInDb         int64   `json:"idInDb"`
	UploadDuration float64 `json:"uploadDuration"`
}
type Creation_ER struct {
	*Creation
	Error string `json:"error"`
}

type CreationInfo_OK struct {
	CreationId int64   `json:"vid"`
	AuthorId   int64   `json:"uid"`
	Title      string  `json:"title"`
	Duration   float64 `json:"duration"`
}

type CreationInfo_ER struct {
	CreationId int64  `json:"vid"`
	AuthorId   int64  `json:"uid"`
	Title      string `json:"title"`
	Error      string `json:"error"`
}
