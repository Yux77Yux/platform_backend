package model

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
