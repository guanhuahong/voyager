package club

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	clubReq "github.com/flipped-aurora/gin-vue-admin/server/model/club/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TbClubListService struct {
}

// CreateTbClubList 创建TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) CreateTbClubList(tbClubList club.TbClubList) (err error) {
	err = global.GVA_DB.Create(&tbClubList).Error
	return err
}

// DeleteTbClubList 删除TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) DeleteTbClubList(tbClubList club.TbClubList) (err error) {
	err = global.GVA_DB.Delete(&tbClubList).Error
	return err
}

// DeleteTbClubListByIds 批量删除TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) DeleteTbClubListByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]club.TbClubList{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTbClubList 更新TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) UpdateTbClubList(tbClubList club.TbClubList) (err error) {
	err = global.GVA_DB.Save(&tbClubList).Error
	return err
}

// GetTbClubList 根据id获取TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) GetTbClubList(id uint) (tbClubList club.TbClubList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tbClubList).Error
	return
}

// GetTbClubListInfoList 分页获取TbClubList记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubListService *TbClubListService) GetTbClubListInfoList(info clubReq.TbClubListSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&club.TbClubList{})
	var tbClubLists []club.TbClubList
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.OwnerId != nil {
		db = db.Where("owner_id = ?", info.OwnerId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tbClubLists).Error
	return tbClubLists, total, err
}

func (tbClubListService *TbClubListService) GetClubOptions() (list interface{}, err error) {
	db := global.GVA_DB.Model(&club.TbClubList{})
	var tbClubOptions []club.TbClubOption
	err = db.Find(&tbClubOptions).Error

	return tbClubOptions, err
}
