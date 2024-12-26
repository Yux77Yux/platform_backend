package service

import "log"

var (
	addr      string
	http_addr string
)

func InitStr(Addr string, HttpAddr string) {
	log.Printf(" addr %s", Addr)
	addr = Addr
	http_addr = HttpAddr
}
