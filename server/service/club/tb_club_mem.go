package club

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	clubReq "github.com/flipped-aurora/gin-vue-admin/server/model/club/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

	sqlHelper "github.com/flipped-aurora/gin-vue-admin/server/utils/sql"
)

type TbClubMemService struct {
}

// CreateTbClubMem 创建TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) CreateTbClubMem(tbClubMem club.TbClubMem) (err error) {
	err = global.GVA_DB.Create(&tbClubMem).Error
	return err
}

// DeleteTbClubMem 删除TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) DeleteTbClubMem(tbClubMem club.TbClubMem) (err error) {
	err = global.GVA_DB.Delete(&tbClubMem).Error
	return err
}

// DeleteTbClubMemByIds 批量删除TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) DeleteTbClubMemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]club.TbClubMem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTbClubMem 更新TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) UpdateTbClubMem(tbClubMem club.TbClubMem) (err error) {
	err = global.GVA_DB.Save(&tbClubMem).Error
	return err
}

// GetTbClubMem 根据id获取TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) GetTbClubMem(id uint) (tbClubMem club.TbClubMem, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tbClubMem).Error
	return
}

type TbClubMemListResult struct {
	club.TbClubMem
	ClubName string `json:"clubName" form:"clubName" gorm:"column:club_name;"`
}

// GetTbClubMemInfoList 分页获取TbClubMem记录
// Author [piexlmax](https://github.com/piexlmax)
func (tbClubMemService *TbClubMemService) GetTbClubMemInfoList(info clubReq.TbClubMemSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	// db := global.GVA_DB.Model(&club.TbClubMem{})
	db := global.GVA_DB.Table(new(club.TbClubMem).TableName() + " main_table")
	var tbClubMems []TbClubMemListResult

	db = db.Select("main_table.*, c.name as club_name").
		Joins("Left Join tb_club_list c on c.id = main_table.club_id")
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ClubId != nil {
		db = db.Where("main_table.club_id = ?", info.ClubId)
	}
	if info.PlayerId != nil {
		db = db.Where("main_table.player_id = ?", info.PlayerId)
	}
	if info.MemType != nil {
		db = db.Where("main_table.mem_type = ?", info.MemType)
	}
	if info.ParentId != nil {
		db = db.Where("main_table.parent_id = ?", info.ParentId)
	}
	if info.InviteCode != "" {
		db = db.Where("main_table.invite_code = ?", info.InviteCode)
	}

	db = sqlHelper.DbWhereSearchDateRange(db, info.SearchDateRangeInfo, "main_table")

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&tbClubMems).Error
	return tbClubMems, total, err
}
