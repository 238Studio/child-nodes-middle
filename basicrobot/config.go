package basicrobot

import (
	"encoding/json"
	"strconv"
	"time"

	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-config-service/configService"
	"github.com/UniversalRobotDriveTeam/child-nodes-database-service/databaseService"
	"github.com/UniversalRobotDriveTeam/child-nodes-device-service/deviceService"
	"github.com/UniversalRobotDriveTeam/child-nodes-websocket-service/websocketService"
)

// Config本身没有配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cConfig() {
	return
}

// 根据配置启动Config
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) iConfig() {
	configService0 := configService.InitConfigManager()
	manager.RobotBasicModulesSet.ConfigService = configService0
}

// 根据配置填充本地数据库的初始化参数
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cLocalDatabase() {
	err, databaseURL := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.LocalDatabaseConfig, _const.LocalDatabaseURL)
	if err != nil {
		//todo
	}
	err, databaseName := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.LocalDatabaseConfig, _const.LocalDatabaseName)
	var value interface{}
	value = &[]string{databaseURL, databaseName}
	manager.BasicRobotModulesInitArgsSet.LocalDatabaseServiceInit = value
}

// 根据配置初始化本地服务器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) iLocalDatabase() {
	args := (manager.BasicRobotModulesInitArgsSet.LocalDatabaseServiceInit).(*[]string)
	databaseName0 := (*args)[0]
	databaseURL0 := (*args)[1]
	localDatabaseService, err := databaseService.InitSQLiteDatabase(databaseName0, databaseURL0)
	if err != nil {

	}
	//todo:err
	manager.RobotBasicModulesSet.LocalDatabaseService = localDatabaseService
}

// 根据配置填充远程数据库的初始化参数
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cRemoteDatabase() {
	err, databaseURL := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.RemoteDatabaseConfig, _const.RemoteDatabaseURL)
	if err != nil {
		//todo
	}
	err, databaseName := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.RemoteDatabaseConfig, _const.RemoteDatabaseName)
	var value interface{}
	value = &[]string{databaseURL, databaseName}
	manager.BasicRobotModulesInitArgsSet.RemoteDatabaseServiceInit = value
}

// 根据配置初始化远程服务器
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) iRemoteDatabase() {
	args := (manager.BasicRobotModulesInitArgsSet.RemoteDatabaseServiceInit).(*[]string)
	databaseName0 := (*args)[0]
	databaseURL0 := (*args)[1]
	remoteDatabaseService, err := databaseService.InitSQLiteDatabase(databaseName0, databaseURL0)
	if err != nil {

	}
	//todo:err
	manager.RobotBasicModulesSet.RemoteDatabaseService = remoteDatabaseService
}

// 根据传入配置初始化串口配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cSerial() {
	err, Baud_ := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.SerialConfig, _const.SerialBaud)
	if err != nil {

	}
	//todo:err
	Baud, err := strconv.ParseInt(Baud_, 10, 64)
	err, readTimeOut_ := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.SerialConfig, _const.SerialReadTimeOutMil)
	readTimeOut, err := strconv.ParseInt(readTimeOut_, 10, 64)
	err, sendCacheLength_ := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.SerialConfig, _const.SerialSendCacheLength)
	args := make([]interface{}, 3)
	args[0] = Baud
	args[1] = readTimeOut
	args[2], _ = strconv.ParseInt(sendCacheLength_, 10, 64)
	manager.BasicRobotModulesInitArgsSet.SerialServiceInit = &args
}

// 初始化串口
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) iSerial() {
	Baud := 0
	readTimeOut := time.Millisecond
	args := (manager.BasicRobotModulesInitArgsSet.SerialServiceInit).([]interface{})
	Baud = args[0].(int)
	readTimeOut = time.Duration(args[1].(int)) * time.Millisecond
	sendCacheLength := args[2].(int)
	manager.RobotBasicModulesSet.SerialService = deviceService.InitSerialApp(Baud, readTimeOut, sendCacheLength)
}

// 初始化错误管理器的配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cErrorManager() {
	if manager.RobotBasicModulesSet.LocalDatabaseService == nil {
		return
		//todo
	}
	if !manager.RobotBasicModulesSet.LocalDatabaseService.IsAlive() {
		return
		//todo
	}
	var value interface{}
	value = manager.RobotBasicModulesSet.LocalDatabaseService
	manager.BasicRobotModulesInitArgsSet.ErrorManagerInit = value
}

// 初始化本地调用配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cNativeCall() {
	url, _ := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.NativeConfig, _const.NativeCallURL)
	manager.BasicRobotModulesInitArgsSet.LocalCallHttpInit = &url
}

// 初始化Http网络服务配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cHttpApp() {
	_, url := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.HTTPConfig, _const.HTTPUrl)
	_, params_ := manager.RobotBasicModulesSet.ConfigService.ReadConfig(_const.HTTPConfig, _const.HTTPParams)
	var params = make(map[string]string)
	err := json.Unmarshal([]byte(params_), &params)
	if err != nil {

	}
	//todo
	value := make([]interface{}, 2)
	value[0] = url
	value[1] = params
	manager.BasicRobotModulesInitArgsSet.HttpAppInit = &value
}

// 初始化websocket网络服务配置
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) cWebApp() {
	var wsURL string
	var pingTimerMargin int
	var pongMaxDelay int
	var value interface{}
	value = []interface{}{
		wsURL, pingTimerMargin, pongMaxDelay,
	}
	manager.BasicRobotModulesInitArgsSet.WebAppInit = &value
}

// 初始化websocket网络服务
// 传入：无
// 传出：无
func (manager *BasicRobotModuleManager) iWebApp() {
	args := *(manager.BasicRobotModulesInitArgsSet.WebAppInit).(*[]interface{})
	wsURL := args[0].(string)
	pingTimerMargin := args[1].(int)
	pongMaxDelay := args[2].(int)
	manager.RobotBasicModulesSet.WebApp, _ = websocketService.InitWebsocketAPP(wsURL, pingTimerMargin, pongMaxDelay)
}
