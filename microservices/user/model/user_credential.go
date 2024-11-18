package model

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"` // 可选邮箱，方便密码重置
}
