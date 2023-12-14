package call

// Call 发起调用
// 传入:包名、函数名、参数
// 传出:返回值
func Call(packageName string, methodName string, args uintptr) uintptr {
	// 获得调用者
	caller := manager.store[packageName]

	// 创建调用表单
	callForm := new(CallForm)
	callForm.packageName = packageName
	callForm.methodName = methodName
	callForm.args = args
	callForm.re = make(chan uintptr)

	// 传递调用信息
	caller.GetCall <- callForm
	return <-callForm.re
}

// GetMethodName 获取函数名
// 传入:无
// 传出:函数名
func (callForm *CallForm) GetMethodName() string {
	return callForm.methodName
}

// GetArgs 获取参数
// 传入:无
// 传出:参数
func (callForm *CallForm) GetArgs() uintptr {
	return callForm.args
}

// Return 返回值
// 传入:返回值
// 传出:无
func (callForm *CallForm) Return(re uintptr) {
	callForm.re <- re
}
