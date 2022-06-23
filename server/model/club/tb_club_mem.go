// 自动生成模板TbClubMem
package club

import (
	"time"
)

// TbClubMem 结构体
// 如果含有time.Time 请自行import time包
type TbClubMem struct {
	ID         uint      `gorm:"primarykey"` // 主键ID
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 更新时间
	ClubId     *int      `json:"clubId" form:"clubId" gorm:"column:club_id;comment:俱乐部id;size:19;"`
	PlayerId   *int      `json:"playerId" form:"playerId" gorm:"column:player_id;comment:成员id;size:19;"`
	MemType    *int      `json:"memType" form:"memType" gorm:"column:mem_type;comment:成员类型;size:10;"`
	ParentId   *int      `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:上级合伙人;size:19;"`
	CutRate    *float64  `json:"cutRate" form:"cutRate" gorm:"column:cut_rate;comment:分成比例;size:19;"`
	InviteCode string    `json:"inviteCode" form:"inviteCode" gorm:"column:invite_code;comment:邀请码;size:20;"`
	PartnerMsg string    `json:"partnerMsg" form:"partnerMsg" gorm:"column:partner_msg;comment:合伙人公告;size:256;"`
}

// TableName TbClubMem 表名
func (TbClubMem) TableName() string {
	return "tb_club_mem"
}
