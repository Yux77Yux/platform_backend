package errMap

import (
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapMySQLErrorToStatus(err error) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062:
			return status.Errorf(codes.AlreadyExists, "Duplicate entry: %s", mysqlErr.Message)
		case 1452:
			return status.Errorf(codes.FailedPrecondition, "Foreign key constraint failed: %s", mysqlErr.Message)
		case 1048:
			return status.Errorf(codes.InvalidArgument, "Column cannot be null: %s", mysqlErr.Message)
		case 1146:
			return status.Errorf(codes.NotFound, "Table not found: %s", mysqlErr.Message)
		case 1054:
			return status.Errorf(codes.InvalidArgument, "Unknown column: %s", mysqlErr.Message)
		case 1366:
			return status.Errorf(codes.InvalidArgument, "Incorrect value: %s", mysqlErr.Message)
		case 2006, 2013:
			return status.Errorf(codes.Unavailable, "Database connection lost: %s", mysqlErr.Message)
		case 1205:
			return status.Errorf(codes.DeadlineExceeded, "Lock wait timeout exceeded: %s", mysqlErr.Message)
		case 1213:
			return status.Errorf(codes.Aborted, "Deadlock found: %s", mysqlErr.Message)
		case 1045:
			return status.Errorf(codes.PermissionDenied, "Access denied: %s", mysqlErr.Message)
		default:
			return status.Errorf(codes.Internal, "Database error: %s", mysqlErr.Message)
		}
	}
	return status.Errorf(codes.Internal, "Unknown error: %s", err.Error())
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
