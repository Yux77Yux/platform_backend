package internal

var (
	messaging MessageQueueMethod
)

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}
