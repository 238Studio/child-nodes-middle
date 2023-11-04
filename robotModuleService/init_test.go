package robotModuleService

import "testing"

func TestName(t *testing.T) {
	manager := InitRobotModuleManagerApp()
	manager.InitAllBasicModels()
}
