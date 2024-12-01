package internal

var (
	// 需要针对每个exchange进行监听处理，统一启动进程存疑
	// userProcessChannel chan func(reqID string) error
	userRequestChannel chan func(reqID string) error
)

func EmpowerDispatch(master DispatcherInterface) {
	userRequestChannel = master.GetChannel()
}

// func EmpowerProcess(master ProcessorInterface) {
// 	userProcessChannel = master.GetChannel()
// }
