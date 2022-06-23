// 自动生成模板TbClubList
package club

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TbClubList 结构体
// 如果含有time.Time 请自行import time包
type TbClubList struct {
	global.GVA_MODEL
	Name    string   `json:"name" form:"name" gorm:"column:name;comment:俱乐部名称;size:64;"`
	OwnerId *int     `json:"ownerId" form:"ownerId" gorm:"column:owner_id;comment:创始人id;size:19;"`
	ClubMsg string   `json:"clubMsg" form:"clubMsg" gorm:"column:club_msg;comment:俱乐部公告;size:511;"`
	FeeSum  *float64 `json:"feeSum" form:"feeSum" gorm:"column:fee_sum;comment:累计服务费;size:19;"`
	CutSum  *float64 `json:"cutSum" form:"cutSum" gorm:"column:cut_sum;comment:累计分成;size:19;"`
}

// TableName TbClubList 表名
func (TbClubList) TableName() string {
	return "tb_club_list"
}

type TbClubOption struct {
	ID   uint   `gorm:"primarykey"` // 主键ID
	Name string `json:"name" form:"name" gorm:"column:name;"`
}
