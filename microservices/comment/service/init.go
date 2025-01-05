package service

import "log"

var (
	addr string
)

func InitStr(Addr string) {
	log.Printf(" addr %s", Addr)
	addr = Addr
}
