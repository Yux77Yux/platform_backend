package messaging

const (
	Register         = "Register"
	StoreUser        = "StoreUser"
	StoreCredentials = "StoreCredentials"
	UpdateUserSpace  = "UpdateUserSpace"
	UpdateUserBio    = "UpdateUserBio"
	UpdateUserAvatar = "UpdateUserAvatar"
	Follow           = "Follow"

	// review
	UpdateUserStatus = "UpdateUserStatus" // 终点
	DelReviewer      = "DelReviewer"      // 终点
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		Register:         "direct",
		StoreUser:        "direct",
		StoreCredentials: "direct",
		UpdateUserSpace:  "direct",
		UpdateUserBio:    "direct",
		UpdateUserAvatar: "direct",
		UpdateUserStatus: "direct",
		DelReviewer:      "direct",
		Follow:           "direct",
	}
)

func InitStr(_str string) {
	connStr = _str
}
