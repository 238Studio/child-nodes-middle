package errorManager

import "github.com/UniversalRobotDriveTeam/child-nodes-assist/util"

// RecordError 记录错误 将错误写入数据库
// 传入：错误
// 传出：无
func (manager *ErrorManager) RecordError(customError *util.CustomError) {
	err := manager.database.AddData(manager.errorDatabaseTable, customError)
	if err != nil {
		return
	}
}

// QueryErrorsByClass 根据错误等级查询所有时段内存在的错误
// 输入：错误等级，模块
// 输出：符合条件的错误集合
func (manager *ErrorManager) QueryErrorsByClass(errorClass int) *[]util.CustomError {
	r := make([]util.CustomError, 0)
	manager.database.GetDatabase().Table(manager.errorDatabaseTable).Where("errorclass = ?", errorClass).Find(&r)
	return &r
}

// QueryErrorsByClassAndTime 根据错误等级查询一定时间内存在的错误
// 输入：错误等级，开始时间，结束时间(时间戳)，模块
// 输出：符合条件的错误集合
func (manager *ErrorManager) QueryErrorsByClassAndTime(errorClass int, beginTime string, endTime string) *[]util.CustomError {
	r := make([]util.CustomError, 0)
	manager.database.GetDatabase().Table(manager.errorDatabaseTable).Where("errorclass = ? AND date > beginTime AND date < endTime", errorClass, beginTime, endTime).Find(&r)
	return &r
}

// GetErrorChannel 获取向错误管理器传递错误的管道
// 输入：模块名
// 输出：错误消息管道
func (manager *ErrorManager) GetErrorChannel(moduleID string) *ErrorMessageChannel {
	r := ErrorMessageChannel{
		moduleID:     moduleID,
		channelID:    0,
		errorChannel: make(chan *util.CustomError, 1),
	}
	i := -1
A:
	for {
		i++
		for j := range manager.errorChannels[moduleID] {
			if j == i {
				continue A
			}
		}
		break
	}
	r.channelID = i
	return &r
}

// ReleaseErrorChannel 释放错误管道
// 输入：错误管道指针
// 输出：无
func (manager *ErrorManager) ReleaseErrorChannel(channel *ErrorMessageChannel) {
	delete(manager.errorChannels[channel.moduleID], channel.channelID)
	if len(manager.errorChannels[channel.moduleID]) == 0 {
		delete(manager.errorChannels, channel.moduleID)
	}
}

// StartErrorRevFromChannels 启动管道记录错误
// 输入：需要启动的管道错误记录
// 输出：无
func (manager *ErrorManager) StartErrorRevFromChannels(channel *ErrorMessageChannel) {
	go manager.errorRevFromChannels(channel)
}

// StopAllErrorRevFromChannels 关闭全部管道记录错误
// 输入：无
// 输出：无
func (manager *ErrorManager) StopAllErrorRevFromChannels() {
	manager.stopErrorChannelsRev <- struct{}{}
}

// 管道错误记录器
// 输入：无
// 输出：无
func (manager *ErrorManager) errorRevFromChannels(channel *ErrorMessageChannel) {
	for {
		select {
		case <-manager.stopErrorChannelsRev:
			break
		case err := <-channel.errorChannel:
			manager.RecordError(err)
		}
	}
}
