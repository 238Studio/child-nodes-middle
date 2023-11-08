package api

import "github.com/UniversalRobotDriveTeam/child-nodes-reality-objects/realityObjects/space"

// SetExpectSpeed 设置机器人目标速度
// 传入：目标速度
// 传出：无
func (basicTaskAPIApp *BasicTaskApp) SetExpectSpeed(speed space.D3Vector) {

}

// SetExpectAcceleratedSpeed 设置目标加速度大小
// 传入：目标加速度
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExpectAcceleratedSpeed(acceleratedSpeed space.D3Vector) {

}

// SetExpectRotationSpeed 设置目标转动速度大小
// 传入：目标转动速度
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExpectRotationSpeed(rotationSpeed space.D3Vector) {

}

// SetExpectRotationAcceleratedSpeed 设置目标转动加速度大小
// 传入：目标转动加速度
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExpectRotationAcceleratedSpeed(acceleratedRotationSpeed space.D3Vector) {

}

// SetExpectRotationDiff 设置机器人目标转动角度差
// 传入：目标转动角度差 顺时针为正 角度值
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExpectRotationDiff(rotationDiff space.D3Vector) {

}

// SetExpectMobileManipulatorDisplacement 将机械臂移动一个向量
// 传入：机械臂ID，需要移动的向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExpectMobileManipulatorDisplacement(mobileManipulatorID string, v space.D3Vector) {

}

// SetExceptMobileManipulatorSpeed 将机械臂赋予一个速度向量
// 传入：机械臂ID，速度向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExceptMobileManipulatorSpeed(mobileManipulatorID string, v space.D3Vector) {

}

// SetExceptMobileManipulatorAcceleratedSpeed 设置机械臂加速度向量
// 传入：机械臂ID，加速度向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExceptMobileManipulatorAcceleratedSpeed(mobileManipulatorID string, v space.D3Vector) {

}

// SetExceptTripodHeadDisplacement 将云台转动一个向量
// 传入：云台ID，需要转动的向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExceptTripodHeadDisplacement(tripodHeadID string, v space.D3Vector) {

}

// SetExceptTripodHeadSpeed 将云台赋予一个角速度向量
// 传入：云台ID，角速度向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExceptTripodHeadSpeed(tripodHeadID string, v space.D3Vector) {

}

// SetExceptTripodHeadAcceleratedSpeed 设置机械臂加速度向量
// 传入：云台ID，角加速度向量
// 传出：无
func (BasicTaskApp *BasicTaskApp) SetExceptTripodHeadAcceleratedSpeed(tripodHeadID string, v space.D3Vector) {

}
