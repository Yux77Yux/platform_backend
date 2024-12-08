package internal

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ParseTimestamp(field string) (*timestamppb.Timestamp, error) {
	if field == "none" {
		return nil, nil
	}

	result, err := time.Parse(time.RFC3339, field)
	if err != nil {
		return nil, fmt.Errorf("invalid format: %v", err)
	}
	return timestamppb.New(result), nil
}
