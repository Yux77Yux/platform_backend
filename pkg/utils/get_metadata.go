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
		return fmt.Sprintf("warning: metadata %s not found in context", key)
	}

	values := md.Get(key)
	if len(values) == 0 {
		return ""
	}

	return strings.TrimSpace(strings.Split(values[0], ",")[0])
}
