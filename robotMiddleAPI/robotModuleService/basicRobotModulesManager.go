package robotModuleService

import "time"

// StartRobotModulesMonitor 开始子节点模块监控器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) StartRobotModulesMonitor() {
	manager.robotMonitorIsRun = true
	go manager.robotModulesMonitor()
}

// StopRobotModulesMonitor 停止子节点模块监控器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) StopRobotModulesMonitor() {
	*manager.stopRobotMonitorChannel <- 0
	manager.robotMonitorIsRun = false
}

// 子节点监控器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) robotModulesMonitor() {
	for {
		time.Sleep(time.Millisecond * 50)
		select {
		case <-*manager.stopRobotMonitorChannel:
			break
		default:
			if manager.RobotBasicModulesExpectRunSet.ConfigServiceExpectRun && !manager.RobotBasicModulesSet.ConfigService.IsAlive() {
				go manager.RobotBasicModulesSet.ConfigService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.LocalDatabaseServiceExpectRun && !manager.RobotBasicModulesSet.LocalDatabaseService.IsAlive() {
				go manager.RobotBasicModulesSet.LocalDatabaseService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.RemoteDatabaseServiceExpectRun && !manager.RobotBasicModulesSet.RemoteDatabaseService.IsAlive() {
				go manager.RobotBasicModulesSet.RemoteDatabaseService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.SerialServiceExpectRun && !manager.RobotBasicModulesSet.SerialService.IsAlive() {
				go manager.RobotBasicModulesSet.SerialService.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.LocalCallHttpExpectRun && !manager.RobotBasicModulesSet.LocalCallHttp.IsAlive() {
				go manager.RobotBasicModulesSet.LocalCallHttp.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.HttpAppExpectRun && !manager.RobotBasicModulesSet.HttpApp.IsAlive() {
				go manager.RobotBasicModulesSet.HttpApp.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.WebAppExpectRun && !manager.RobotBasicModulesSet.WebApp.IsAlive() {
				go manager.RobotBasicModulesSet.WebApp.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.TaskTableAppExpectRun && !manager.RobotBasicModulesSet.TaskTableApp.IsAlive() {
				go manager.RobotBasicModulesSet.TaskTableApp.Start()
			}
			if manager.RobotBasicModulesExpectRunSet.VisualAppExpectRun && !manager.RobotBasicModulesSet.VisualApp.IsAlive() {
				go manager.RobotBasicModulesSet.VisualApp.Start()
			}

		}
	}
}
