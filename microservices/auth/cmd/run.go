package main

import (
	"os"

	service "github.com/Yux77Yux/platform_backend/microservices/auth/service"
	tools "github.com/Yux77Yux/platform_backend/microservices/auth/tools"
)

func Run(signal chan os.Signal) {
	var (
		closeServer func(chan any)
	)

	closeServer = service.ServerRun()

	<-signal
	s_closed_sig := make(chan any, 1)
	closeServer(s_closed_sig)

	tools.LogInfo("main", "exit")
	os.Exit(0)
}
