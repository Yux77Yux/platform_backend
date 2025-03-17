package dispatch

var (
	messaging MessageQueueMethod
)

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}
