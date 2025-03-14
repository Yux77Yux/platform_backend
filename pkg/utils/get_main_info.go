package utils

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func GetMainValue(ctx context.Context) string {
	id, ok := ctx.Value("main").(uuid.UUID)
	if !ok {
		LogError("TRACE_ID_NOT_FOUND", "RabbitMQClient.Close", fmt.Errorf("trace ID not set in main function"))
		return ""
	}
	return id.String()
}
