package errMap

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapMySQLErrorToStatus(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return status.Errorf(codes.DeadlineExceeded, "Operation timed out: %s", err.Error())
	}

	switch err {
	case sql.ErrNoRows:
		return status.Errorf(codes.NotFound, "No records found")
	case sql.ErrConnDone:
		return status.Errorf(codes.Unavailable, "Database connection closed")
	case sql.ErrTxDone:
		return status.Errorf(codes.FailedPrecondition, "Transaction has already been committed or rolled back")
	}

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
