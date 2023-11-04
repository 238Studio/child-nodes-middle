package basicTaskAPI

import (
	"github.com/UniversalRobotDriveTeam/child-nodes-reality-objects/realityObjects/visualObject"
	"github.com/UniversalRobotDriveTeam/child-nodes-visual-service/visualService/visualServiceStructs"
)

// GetObjectsPhotosInPhoto 使用指定模型切割画面
// 传入：模型集合，图像
// 传出：物品图像集合
func (BasicTaskApp *BasicTaskApp) GetObjectsPhotosInPhoto(visualModules map[string]visualServiceStructs.PictureSegmentationModule, photo visualObject.Photo) {
}

// GetObjectState 获取物品状态，将物品的图像输入，映射到有限个状态
// 传入：分类模型，物品
// 传出：状态
func (BasicTaskApp *BasicTaskApp) GetObjectState(visualModules map[string]visualServiceStructs.VisualClassifyModule, object visualObject.Object) {

}

// GetObjectTags 使用指定模型识别到物品tags
// 传入：模型集合，物品图像集合
// 传出：物品
func (BasicTaskApp *BasicTaskApp) GetObjectTags(visualModules map[string]visualServiceStructs.VisualRecognitionModule, object visualObject.Object) {

}
