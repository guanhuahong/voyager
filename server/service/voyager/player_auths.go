package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
)

type PlayerAuthService struct {
}

// CreatePlayer 创建Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerAuthService *PlayerAuthService) CreatePlayer(player voyager.PlayerAuth) (err error) {
	err = global.GVA_DB.Create(&player).Error
	return err
}

// DeletePlayer 删除Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerAuthService *PlayerAuthService) DeletePlayer(player voyager.PlayerAuth) (err error) {
	err = global.GVA_DB.Delete(&player).Error
	return err
}

// DeletePlayerByIds 批量删除Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerAuthService *PlayerAuthService) DeletePlayerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]voyager.PlayerAuth{}, "player_id in ?", ids.Ids).Error
	return err
}

// UpdatePlayer 更新Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerAuthService *PlayerAuthService) UpdatePlayer(player voyager.PlayerAuth) (err error) {
	err = global.GVA_DB.Save(&player).Error
	return err
}

// GetPlayer 根据id获取Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerAuthService *PlayerAuthService) GetPlayer(id uint) (player voyager.PlayerAuth, err error) {
	err = global.GVA_DB.Where("player_id = ?", id).First(&player).Error
	return
}
