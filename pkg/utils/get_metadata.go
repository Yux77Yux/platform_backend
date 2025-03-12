package utils

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

func GetMetadataValue(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Sprintf("warning: metadata not found in context")
	}

	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}

	return strings.TrimSpace(strings.Split(values[0], ",")[0])
}
