package internal

import (
	userMQ "github.com/Yux77Yux/platform_backend/microservices/user/messaging"
)

var (
	userRequestChannel chan userMQ.RequestHandlerFunc
)

func EmpowerDispatch(master *userMQ.RequestDispatcher) {
	userRequestChannel = master.GetChannel()
}
