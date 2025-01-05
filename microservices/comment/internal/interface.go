package internal

type DispatcherInterface interface {
	GetChannel() chan func(string) error
}

type ProcessorInterface interface {
	GetChannel() chan func(string)
}
