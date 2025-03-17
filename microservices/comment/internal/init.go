package internal

var (
	db        SqlMethod
	messaging MessageQueueMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}

func InitMQ(_messaging MessageQueueMethod) {
	messaging = _messaging
}
