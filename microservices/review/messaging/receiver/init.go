package receiver

var (
	db         SqlMethod
	messaging  MessageQueueMethod
	dispatcher DispatchInterface
)

func InitDb(_db SqlMethod) {
	db = _db
}

func Run(_messaging MessageQueueMethod, _dispatch DispatchInterface) {
	dispatcher = _dispatch
	messaging = _messaging
	for exchange := range ExchangesConfig {
		switch exchange {
		case EXCHANGE_NEW_REVIEW:
			go messaging.ListenToQueue(exchange, QUEUE_NEW_REVIEW, KEY_NEW_REVIEW, NewReviewProcessor)
		case EXCHANGE_PEND_CREATION:
			go messaging.ListenToQueue(exchange, QUEUE_PEND_CREATION, KEY_PEND_CREATION, PendingCreationProcessor)
		case EXCHANGE_BATCH_UPDATE:
			go messaging.ListenToQueue(exchange, QUEUE_BATCH_UPDATE, KEY_BATCH_UPDATE, BatchUpdateProcessor)
		case EXCHANGE_UPDATE:
			go messaging.ListenToQueue(exchange, QUEUE_UPDATE, KEY_UPDATE, UpdateProcessor)
		}
	}
}
