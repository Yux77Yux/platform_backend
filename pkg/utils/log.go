package utils

import (
	"fmt"
	"log"
	"time"

	logger "github.com/Yux77Yux/platform_backend/pkg/logger"
)

var (
	logManager *logger.LoggerManager
)

func init() {
	logManager = logger.GetLoggerManager()
}

func LogSuperError(err error) {
	now := time.Now()
	logManager.SharedLog(&logger.LogMessage{
		Level:     logger.SUPER,
		Timestamp: now,
		Message:   err.Error(),
	})
	logManager.Log(&logger.LogFile{
		Path: "./log/super.log",
		LogMessage: &logger.LogMessage{
			Level:     logger.SUPER,
			Timestamp: now,
			Message:   err.Error(),
		},
	})

	log.Fatalf("super error: %s", err.Error())
}

func LogInfo(traceId, fullName string) {
	domainName, methodName := logManager.SplitFullName(fullName)
	now := time.Now()
	go logManager.SharedLog(&logger.LogMessage{
		Level:     logger.INFO,
		TraceId:   traceId,
		Timestamp: now,
		Position:  fullName,
	})
	go logManager.Log(&logger.LogFile{
		Path: fmt.Sprintf("./log/%s.log", domainName),
		LogMessage: &logger.LogMessage{
			Level:     logger.INFO,
			TraceId:   traceId,
			Timestamp: now,
			Position:  methodName,
		},
	})
}

func LogError(traceId, fullName, message string) {
	domainName, methodName := logManager.SplitFullName(fullName)
	now := time.Now()
	go logManager.SharedLog(&logger.LogMessage{
		Level:     logger.ERROR,
		TraceId:   traceId,
		Timestamp: now,
		Position:  fullName,
		Message:   message,
	})
	go logManager.Log(&logger.LogFile{
		Path: fmt.Sprintf("./log/%s.error.log", domainName),
		LogMessage: &logger.LogMessage{
			Level:     logger.ERROR,
			TraceId:   traceId,
			Timestamp: now,
			Position:  methodName,
			Message:   message,
		},
	})
}
