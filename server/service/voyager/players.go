package voyager

import (
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	voyagerReq "github.com/flipped-aurora/gin-vue-admin/server/model/voyager/request"
	"gorm.io/gorm"
)

type PlayerService struct {
}

// CreatePlayer 创建Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) CreatePlayer(player voyager.PlayerCreated) (err error) {
	sql := global.GVA_DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&player)
	})
	log.Println(sql)
	err = global.GVA_DB.Create(&player).Error
	return err
}

// DeletePlayer 删除Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) DeletePlayer(player voyager.Player) (err error) {
	err = global.GVA_DB.Delete(&player).Error
	return err
}

// DeletePlayerByIds 批量删除Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) DeletePlayerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]voyager.Player{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePlayer 更新Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) UpdatePlayer(player voyager.Player) (err error) {
	err = global.GVA_DB.Save(&player).Error
	return err
}

// GetPlayer 根据id获取Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) GetPlayer(id uint) (player voyager.Player, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&player).Error
	return
}

// GetPlayerInfoList 分页获取Player记录
// Author [piexlmax](https://github.com/piexlmax)
func (playerService *PlayerService) GetPlayerInfoList(info voyagerReq.PlayerSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&voyager.Player{})
	var players []voyager.Player
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Mobile != "" {
		db = db.Where("mobile = ?", info.Mobile)
	}
	if info.Age != nil {
		db = db.Where("age = ?", info.Age)
	}
	if info.Idno != "" {
		db = db.Where("idno LIKE ?", "%"+info.Idno+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&players).Error
	return players, total, err
}
