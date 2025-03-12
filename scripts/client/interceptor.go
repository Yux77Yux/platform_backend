package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// gRPC 拦截器，请求自动加 TraceId
func TraceIDInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	traceID := uuid.New().String()
	ctx = metadata.AppendToOutgoingContext(ctx, "TraceId", traceID)

	// 打印方便调试
	fmt.Printf("Calling %s with TraceId: %s\n", method, traceID)

	// 执行真正的调用
	return invoker(ctx, method, req, reply, cc, opts...)
}
