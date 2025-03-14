package middlewares

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	common "github.com/Yux77Yux/platform_backend/generated/common"
	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
)

type TestCoverageInfo struct {
	TraceID        string
	MethodName     string
	BranchCoverage map[string]bool // 各个分支是否被执行（比如 "if_user_exists" -> true）
	ExecutionPath  []string        // 执行路径的顺序（比如 ["start", "if_user_exists", "update_db", "end"]）
	ExecutionCount map[string]int  // 每个分支、方法执行的次数
	Errors         []string        // 如果失败，记录错误信息
}

// 日志拦截器
func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		traceId := utils.GetMetadataValue(ctx, "trace-id")
		if traceId == "" {
			traceId = utils.GetUuidString()
			ctx = metadata.AppendToOutgoingContext(ctx, "trace-id", traceId)
		}
		fullName := info.FullMethod
		ctx = metadata.AppendToOutgoingContext(ctx, "full-name", fullName)

		utils.LogInfo(traceId, fullName)

		resp, err := handler(ctx, req)
		// 返回之后的
		isServerError, detail, c_err := GetMsg(resp, traceId)
		if c_err != nil {
			// 反射的错误,警告
			utils.LogError(traceId, fullName, c_err)
		} else {
			// 没有反射错误，但有业务上的错误
			if isServerError {
				utils.LogError(traceId, fullName, fmt.Errorf(detail))
				return resp, fmt.Errorf(detail)
			}
		}

		// 其他未知错误
		if err != nil {
			utils.LogError(traceId, fullName, err)
			return resp, nil
		}

		utils.LogInfo(traceId, fullName)
		return resp, nil
	}
}

// (ServerError?,ErrorDetail,error)
func GetMsg(response any, traceId string) (bool, string, error) {
	if response == nil {
		return false, "", nil
	}

	v := reflect.ValueOf(response)
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

	v = v.Elem()
	if v.Type() == reflect.TypeOf(&emptypb.Empty{}) {
		return false, "", nil
	}

	msgField := v.FieldByName("Msg")
	if !msgField.IsValid() {
		// Msg 字段不存在（结构体里根本没这个字段）
		log.Printf("msgField %v", response)
		return true, "", fmt.Errorf("error: 未找到 Msg 字段")
	}

	if msgField.Kind() != reflect.Ptr {
		// Msg 字段存在但不是指针类型
		return true, "", fmt.Errorf("error: Msg 字段类型错误，非指针")
	}

	if msgField.IsNil() {
		// Msg 字段是指针，但为空（nil），这种情况不是错误
		return false, "", nil
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
