package errMap

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetStatusError(err error) error {
	switch err {
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "operation timed out")
	case context.Canceled:
		return status.Error(codes.Canceled, "operation was canceled")
	default:
		return status.Error(codes.Unknown, "unknown context error")
	}
}

// GrpcCodeToHTTPStatusString 将 gRPC 错误映射为 HTTP 状态码字符串
func GrpcCodeToHTTPStatusString(err error) string {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.OK:
			return "200"
		case codes.InvalidArgument:
			return "400"
		case codes.NotFound:
			return "404"
		case codes.AlreadyExists:
			return "409"
		case codes.PermissionDenied:
			return "403"
		case codes.Unauthenticated:
			return "401"
		case codes.ResourceExhausted:
			return "429"
		case codes.FailedPrecondition, codes.Aborted:
			return "409"
		case codes.Unavailable:
			return "503"
		case codes.DeadlineExceeded:
			return "504"
		case codes.Internal:
			return "500"
		default:
			return "500"
		}
	}

	// 如果不是 gRPC 错误，统一返回 500
	return "500"
}

func IsServerError(err error) bool {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.Internal,
			codes.Unavailable,
			codes.DeadlineExceeded:
			return true
		}
	}
	// 非 gRPC 错误或者不是服务器端错误，都返回 false
	return false
}
