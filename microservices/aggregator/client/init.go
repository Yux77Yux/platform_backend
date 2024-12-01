package client

var (
	user_service_address string
	auth_service_address string
)

func InitStr(USER_SERVER_ADDRESS, AUTH_SERVER_ADDRESS string) {
	user_service_address = USER_SERVER_ADDRESS
	auth_service_address = AUTH_SERVER_ADDRESS
}
