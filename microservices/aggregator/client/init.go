package client

var (
	service_address string
)

// 使用了envoy，所以使用envoy地址即可
func InitStr(SERVER_ADDRESS string) {
	service_address = SERVER_ADDRESS
}
