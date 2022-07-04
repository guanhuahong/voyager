package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    voyagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/voyager/request"
)

type PlayerExtendService struct {
}

// CreatePlayerExtend 创建PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService) CreatePlayerExtend(playerExtend voyager.PlayerExtend) (err error) {
	err = global.GVA_DB.Create(&playerExtend).Error
	return err
}

// DeletePlayerExtend 删除PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService)DeletePlayerExtend(playerExtend voyager.PlayerExtend) (err error) {
	err = global.GVA_DB.Delete(&playerExtend).Error
	return err
}

// DeletePlayerExtendByIds 批量删除PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService)DeletePlayerExtendByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]voyager.PlayerExtend{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePlayerExtend 更新PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService)UpdatePlayerExtend(playerExtend voyager.PlayerExtend) (err error) {
	err = global.GVA_DB.Save(&playerExtend).Error
	return err
}

// GetPlayerExtend 根据id获取PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService)GetPlayerExtend(id uint) (playerExtend voyager.PlayerExtend, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&playerExtend).Error
	return
}

// GetPlayerExtendInfoList 分页获取PlayerExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerExtendService *PlayerExtendService)GetPlayerExtendInfoList(info voyagerReq.PlayerExtendSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&voyager.PlayerExtend{})
    var playerExtends []voyager.PlayerExtend
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.PlayerId != nil {
        db = db.Where("player_id = ?",info.PlayerId)
    }
    if info.Field != "" {
        db = db.Where("field = ?",info.Field)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&playerExtends).Error
	return  playerExtends, total, err
}
