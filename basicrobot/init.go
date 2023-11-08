package basicrobot

// InitRobotModuleManagerApp 初始化子节点模块管理器
// 传入：初始化参数
// 传出：子节点模块管理器
func InitRobotModuleManagerApp() *BasicRobotModuleManager {
	// 分配基本模块管理器对象
	app := *new(BasicRobotModuleManager)

	// 初始化模块运行状态表
	app.RobotBasicModulesExpectRunSet = new(RobotBasicModulesExpectRun)
	app.RobotBasicModulesExpectRunSet.ConfigServiceExpectRun = true
	app.RobotBasicModulesExpectRunSet.LocalDatabaseServiceExpectRun = true
	app.RobotBasicModulesExpectRunSet.RemoteDatabaseServiceExpectRun = true
	app.RobotBasicModulesExpectRunSet.SerialServiceExpectRun = true
	app.RobotBasicModulesExpectRunSet.ErrorManagerExpectRun = true
	app.RobotBasicModulesExpectRunSet.LocalCallHttpExpectRun = true
	app.RobotBasicModulesExpectRunSet.LocalCallUnixSocketExpectRun = true
	app.RobotBasicModulesExpectRunSet.HttpAppExpectRun = true
	app.RobotBasicModulesExpectRunSet.WebAppExpectRun = true
	app.RobotBasicModulesExpectRunSet.VisualAppExpectRun = true

	// 获取终止模块监控管道
	channel := make(chan int, 1)
	app.stopRobotMonitorChannel = &channel

	//TODO：注释
	app.BasicRobotModulesInitArgsSet = &BasicRobotModulesInitArgs{
		ConfigServiceInit:         nil,
		LocalDatabaseServiceInit:  nil,
		RemoteDatabaseServiceInit: nil,
		SerialServiceInit:         nil,
		ErrorManagerInit:          nil,
		LocalCallHttpInit:         nil,
		LocalCallUnixSocketInit:   nil,
		HttpAppInit:               nil,
		WebAppInit:                nil,
		VisualAppInit:             nil,
	}
	app.RobotBasicModulesSet = &RobotBasicModules{
		ConfigService:         nil,
		LocalDatabaseService:  nil,
		RemoteDatabaseService: nil,
		SerialService:         nil,
		WebApp:                nil,
		//VisualApp:             nil,
	}

	// 调用各个模块的初始化参数载入模块
	app.InitAllBasicModels()
	// 启动基础模块监视器
	app.StartRobotModulesMonitor()
	return &app
}
