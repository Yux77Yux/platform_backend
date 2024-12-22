package client

var (
	user_service_address string
	auth_service_address string
)

func InitStr(SERVER_ADDRESS string) {
	user_service_address = SERVER_ADDRESS
	auth_service_address = SERVER_ADDRESS
}
