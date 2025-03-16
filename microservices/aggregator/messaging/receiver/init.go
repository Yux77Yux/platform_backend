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
		case EXCHANGE_ADD_VIEW:
			go messaging.ListenToQueue(exchange, EXCHANGE_ADD_VIEW, KEY_ADD_VIEW, addViewProcessor)
		}
	}
}
