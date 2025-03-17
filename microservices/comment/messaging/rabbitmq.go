package messaging

import (
	"context"

	internal "github.com/Yux77Yux/platform_backend/microservices/comment/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/comment/messaging/receiver"

	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

func Run(ctx context.Context) func() {
	_client := pkgMQ.GetClient(connStr)

	internal.InitMQ(_client)

	_dispatch := dispatch.Run()
	receiver.Run(_client, _dispatch)

	return func() {
		_client.Close(ctx)
		_dispatch.Close()
	}
}
