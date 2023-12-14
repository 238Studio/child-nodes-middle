package call

var manager *Manager

// InitManager 初始化调用器
// 传入:无
// 传出:无
func InitManager() {
	manager = new(Manager)
	manager.store = make(map[string]*Receiver)
}

// InitReceiver 初始化接收者
// 传入:包名
// 传出:调用者对象
func InitReceiver(packageName string) *Receiver {
	//初始化接收者
	caller := new(Receiver)
	caller.GetCall = make(chan *CallForm)

	//加入管理器
	manager.store[packageName] = caller

	return caller
}
