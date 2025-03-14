package tools

import (
	"context"

	utils "github.com/Yux77Yux/platform_backend/pkg/utils"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CheckStringLength(obj string, min, max int) error {
	return utils.CheckStringLength(obj, min, max)
}

func GetMetadataValue(ctx context.Context, key string) string {
	return utils.GetMetadataValue(ctx, key)
}

func EnsureTimestampPB(input interface{}) (*timestamppb.Timestamp, error) {
	return utils.EnsureTimestampPB(input)
}

func GetUuid() uuid.UUID {
	return utils.GetUuid()
}

func GetUuidString() string {
	return utils.GetUuidString()
}

func LogSuperError(err error) {
	utils.LogSuperError(err)
}
func LogError(traceId, fullName string, err error) {
	utils.LogError(traceId, fullName, err)
}
func LogInfo(traceId, fullName string) {
	utils.LogInfo(traceId, fullName)
}
func LogWarning(traceId, fullName, warning string) {
	utils.LogWarning(traceId, fullName, warning)
}

func GetMainValue(ctx context.Context) string {
	return utils.GetMainValue(ctx)
}
