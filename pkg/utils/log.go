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
	if err == nil {
		return
	}
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

func LogError(traceId, fullName string, err error) {
	if err == nil {
		return
	}
	domainName, methodName := logManager.SplitFullName(fullName)
	now := time.Now()
	go logManager.SharedLog(&logger.LogMessage{
		Level:     logger.ERROR,
		TraceId:   traceId,
		Timestamp: now,
		Position:  fullName,
		Message:   err.Error(),
	})
	go logManager.Log(&logger.LogFile{
		Path: fmt.Sprintf("./log/%s.error.log", domainName),
		LogMessage: &logger.LogMessage{
			Level:     logger.ERROR,
			TraceId:   traceId,
			Timestamp: now,
			Position:  methodName,
			Message:   err.Error(),
		},
	})
}

func LogWarning(traceId, fullName, warning string) {
	domainName, methodName := logManager.SplitFullName(fullName)
	now := time.Now()
	go logManager.SharedLog(&logger.LogMessage{
		Level:     logger.WARNING,
		TraceId:   traceId,
		Timestamp: now,
		Position:  fullName,
		Message:   warning,
	})
	go logManager.Log(&logger.LogFile{
		Path: fmt.Sprintf("./log/%s.warning.log", domainName),
		LogMessage: &logger.LogMessage{
			Level:     logger.WARNING,
			TraceId:   traceId,
			Timestamp: now,
			Position:  methodName,
			Message:   warning,
		},
	})
}
