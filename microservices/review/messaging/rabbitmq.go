package messaging

import (
	"context"

	internal "github.com/Yux77Yux/platform_backend/microservices/review/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/review/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/review/messaging/receiver"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

func Run(ctx context.Context) func() {
	_client := pkgMQ.GetClient(connStr)
	_dispatch := dispatch.Run()
	internal.InitMQ(_client)

	receiver.Run(_client, _dispatch)

	return func() {
		_client.Close(ctx)
		_dispatch.Close()
	}
}
