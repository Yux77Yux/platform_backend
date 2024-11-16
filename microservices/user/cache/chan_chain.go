package cache

// 责任链注册与关闭
type ChanChainMethods interface {
	register(next *ChanChain)
	closeChan() error
}

type ChanChain struct {
	closeChanClosure func() error // 更改为返回 error
	next             *ChanChain
}

func (a *ChanChain) register(next *ChanChain) {
	if a.next != nil {
		// 头插法
		next.next = a.next
		a.next = next
	} else {
		a.next = next
	}
}

func (a *ChanChain) closeChan() error {
	if a.next != nil {
		if err := a.next.closeChan(); err != nil {
			return err
		}
	}
	return a.closeChanClosure() // 现在返回 error
}
