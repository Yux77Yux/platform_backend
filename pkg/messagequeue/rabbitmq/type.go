package rabbitmq

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type HandlerFunc func(context.Context, *anypb.Any) error
type HandlerFuncWithReturn func(context.Context, *anypb.Any) (proto.Message, error)
