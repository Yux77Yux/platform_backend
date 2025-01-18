package internal

type RequestHandlerFunc = func(string) error

type DispatcherInterface interface {
	Start()
	GetChannel() chan RequestHandlerFunc
	Shutdown()
}
