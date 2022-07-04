package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    voyagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/voyager/request"
)

type PlayerProfileService struct {
}

// CreatePlayerProfile 创建PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService) CreatePlayerProfile(playerProfile voyager.PlayerProfile) (err error) {
	err = global.GVA_DB.Create(&playerProfile).Error
	return err
}

// DeletePlayerProfile 删除PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService)DeletePlayerProfile(playerProfile voyager.PlayerProfile) (err error) {
	err = global.GVA_DB.Delete(&playerProfile).Error
	return err
}

// DeletePlayerProfileByIds 批量删除PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService)DeletePlayerProfileByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]voyager.PlayerProfile{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePlayerProfile 更新PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService)UpdatePlayerProfile(playerProfile voyager.PlayerProfile) (err error) {
	err = global.GVA_DB.Save(&playerProfile).Error
	return err
}

// GetPlayerProfile 根据id获取PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService)GetPlayerProfile(id uint) (playerProfile voyager.PlayerProfile, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&playerProfile).Error
	return
}

// GetPlayerProfileInfoList 分页获取PlayerProfile记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerProfileService *PlayerProfileService)GetPlayerProfileInfoList(info voyagerReq.PlayerProfileSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&voyager.PlayerProfile{})
    var playerProfiles []voyager.PlayerProfile
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.PlayerId != nil {
        db = db.Where("player_id = ?",info.PlayerId)
    }
    if info.Nick != "" {
        db = db.Where("nick LIKE ?","%"+ info.Nick+"%")
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&playerProfiles).Error
	return  playerProfiles, total, err
}
