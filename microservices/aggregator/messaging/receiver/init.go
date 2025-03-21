package receiver

var (
	cache      CacheMethod
	dispatcher DispatchInterface
)

func InitCache(_cache CacheMethod) {
	cache = _cache
}

func Run(mq MessageQueueMethod, _dispatch DispatchInterface) {
	messaging := mq
	dispatcher = _dispatch
	for exchange := range ExchangesConfig {
		switch exchange {
		case EXCHANGE_INCREASE_VIEW:
			go messaging.ListenToQueue(exchange, QUEUE_INCREASE_VIEW, KEY_INCREASE_VIEW, increaseViewProcessor)
		}
	}
}
