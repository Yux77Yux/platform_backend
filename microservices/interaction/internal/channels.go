package internal

var (
	RequestChannel chan RequestHandlerFunc
)

func EmpowerDispatch(master DispatcherInterface) {
	RequestChannel = master.GetChannel()
}
