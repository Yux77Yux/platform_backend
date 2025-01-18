package internal

var (
	requestChannel chan RequestHandlerFunc
)

func EmpowerDispatch(master DispatcherInterface) {
	requestChannel = master.GetChannel()
}
