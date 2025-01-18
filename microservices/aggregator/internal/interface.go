package internal

type DispatcherInterface interface {
	Start()
	GetChannel() chan RequestHandlerFunc
	Shutdown()
}

type RequestHandlerFunc = func(string) error
