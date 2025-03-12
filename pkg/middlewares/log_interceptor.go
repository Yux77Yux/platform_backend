package middlewares

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	logger "github.com/Yux77Yux/platform_backend/pkg/logger"
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
)

var (
	logManager *logger.LoggerManager
)

func init() {
	logManager = logger.GetLoggerManager()
}

// 日志拦截器
func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now().Truncate(time.Second)

		traceId := utils.GetMetadataValue(ctx, "trace-id")
		if traceId == "" {
			traceId = uuid.New().String()
			ctx = metadata.AppendToOutgoingContext(ctx, "TraceId", traceId)
		}

		fullName := info.FullMethod
		lastSlash := strings.LastIndex(fullName, "/")
		methodName := fullName[lastSlash+1:]
		go logManager.SharedLog(&logger.LogMessage{
			Level:     logger.INFO,
			TraceId:   traceId,
			Timestamp: start,
			Message:   fmt.Sprintf("%s start", fullName),
		})
		go logManager.Log(&logger.LogFile{
			Path: fmt.Sprintf("./log/%s.log", methodName),
			LogMessage: &logger.LogMessage{
				Level:     logger.INFO,
				TraceId:   traceId,
				Timestamp: start,
			},
		})

		resp, err := handler(ctx, req)
		end := time.Now().Truncate(time.Second)

		isServerError, detail, c_err := GetMsg(req, traceId)
		if c_err != nil {
			Extra := make(map[string]interface{})
			Extra["Detail"] = c_err.Error()
			go logManager.SharedLog(&logger.LogMessage{
				Level:     logger.ERROR,
				TraceId:   traceId,
				Timestamp: end,
				Message:   fmt.Sprintf("%s start", fullName),
				Extra:     Extra,
			})
			go logManager.Log(&logger.LogFile{
				Path: fmt.Sprintf("./log/%s.log", methodName),
				LogMessage: &logger.LogMessage{
					Level:     logger.ERROR,
					TraceId:   traceId,
					Timestamp: end,
					Extra:     Extra,
				},
			})
		} else {
			// 没有反射错误，但有业务上的错误
			if isServerError {
				Extra := make(map[string]interface{})
				Extra["Detail"] = detail
				go logManager.SharedLog(&logger.LogMessage{
					Level:     logger.ERROR,
					TraceId:   traceId,
					Timestamp: end,
					Message:   fmt.Sprintf("%s start", fullName),
					Extra:     Extra,
				})
				go logManager.Log(&logger.LogFile{
					Path: fmt.Sprintf("./log/%s.log", methodName),
					LogMessage: &logger.LogMessage{
						Level:     logger.ERROR,
						TraceId:   traceId,
						Timestamp: end,
						Extra:     Extra,
					},
				})

				return resp, fmt.Errorf(detail)
			}
		}

		// 其他未知错误
		if err != nil {
			Extra := make(map[string]interface{})
			Extra["Detail"] = err.Error()
			go logManager.Log(&logger.LogFile{
				Path: fmt.Sprintf("./log/%s.error.log", methodName),
				LogMessage: &logger.LogMessage{
					Level:     logger.ERROR,
					TraceId:   traceId,
					Timestamp: end,
					Message:   fmt.Sprintf("%s error", methodName),
					Extra:     Extra,
				},
			})

			go logManager.SharedLog(&logger.LogMessage{
				Level:     logger.ERROR,
				TraceId:   traceId,
				Timestamp: end,
				Message:   fmt.Sprintf("%s error", fullName),
				Extra:     Extra,
			})

			return resp, nil
		}

		go logManager.SharedLog(&logger.LogMessage{
			Level:     logger.INFO,
			TraceId:   traceId,
			Timestamp: end,
			Message:   fmt.Sprintf("%s success", fullName),
		})
		go logManager.Log(&logger.LogFile{
			Path: fmt.Sprintf("./log/%s.log", methodName),
			LogMessage: &logger.LogMessage{
				Level:     logger.INFO,
				TraceId:   traceId,
				Timestamp: end,
				Message:   fmt.Sprintf("%s success", methodName),
			},
		})
		return resp, nil
	}
}

// (ServerError?,ErrorDetail,error)
func GetMsg(req any, traceId string) (bool, string, error) {
	v := reflect.ValueOf(req)
	if !v.IsValid() {
		// 不可用
		return false, "", nil
	}

	isNil := v.IsNil()
	if isNil {
		return false, "", nil
	}

	kind := v.Kind()
	if kind != reflect.Pointer {
		return false, "", nil
	}

	elem := v.Elem()
	if elem.Type() == reflect.TypeOf(&emptypb.Empty{}) {
		return false, "", nil
	}

	msgField := v.FieldByName("Msg")
	if !msgField.IsValid() || msgField.IsNil() || msgField.Kind() != reflect.Ptr {
		return true, "", fmt.Errorf("error: Msg 字段无效或为空")
	}

	// 类型断言，确保 Msg 字段是 *common.ApiResponse 类型
	msg, ok := msgField.Interface().(*common.ApiResponse)
	if !ok {
		return true, "", fmt.Errorf("error: Msg 字段类型错误")
	}

	msg.TraceId = traceId
	status := msg.GetStatus()
	code := msg.GetCode()

	if status != common.ApiResponse_PENDING && status != common.ApiResponse_SUCCESS {
		detail := msg.GetDetails()
		if len(code) > 0 && code[0] == '5' {
			return true, detail, nil
		}
		return false, detail, nil
	}

	return false, "", nil
}
