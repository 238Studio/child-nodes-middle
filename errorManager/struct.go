package errorManager

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	"github.com/UniversalRobotDriveTeam/child-nodes-database-service/databaseService"
)

// ErrorManager 错误管理器
type ErrorManager struct {
	// 数据库应用
	database databaseService.DatabaseAPP
	// 错误数据库数据表表名
	errorDatabaseTable string
	// 错误管道集合 moduleID->channelID->*channel
	errorChannels map[string]map[int]*ErrorMessageChannel
	// 停止错误收集管道
	stopErrorChannelsRev chan struct{}
}

// ErrorBasicStatistician 错误基础统计器
type ErrorBasicStatistician struct {
}

// ErrorMessageChannel 错误消息管道
type ErrorMessageChannel struct {
	moduleID string
	// 模块ID
	channelID int
	// 管道ID
	errorChannel chan *util.CustomError
	// 错误管道
}
