package service

var (
	addr      string
	http_addr string
)

func InitStr(Addr string, HttpAddr string) {
	addr = Addr
	http_addr = HttpAddr
}
