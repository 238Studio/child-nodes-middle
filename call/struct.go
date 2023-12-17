package call

// Manager 调用管理器对象
type Manager struct {
	store map[string]*Receiver // 存储调用者对象
}

// Receiver 接收者对象
type Receiver struct {
	GetCall chan *CallForm // 用于传递调用信息
}

// CallForm 调用表单
// 用于传递调用信息。包含包名、函数名、参数等信息
type CallForm struct {
	packageName string       // 包名
	methodName  string       // 函数名
	args        []uintptr    // 参数
	re          chan uintptr // 返回值，通过通道传递
}
