package messaging

import (
	"context"

	internal "github.com/Yux77Yux/platform_backend/microservices/interaction/internal"
	dispatch "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/dispatch"
	receiver "github.com/Yux77Yux/platform_backend/microservices/interaction/messaging/receiver"
	pkgMQ "github.com/Yux77Yux/platform_backend/pkg/messagequeue/rabbitmq"
)

func Run(ctx context.Context) func() {
	_client := pkgMQ.GetClient(connStr)
	internal.InitMQ(_client)
	dispatch.InitMQ(_client)
	_dispatch := dispatch.Run()

	receiver.Run(_client, _dispatch)

	return func() {
		_client.Close(ctx)
		_dispatch.Close()
	}
}
