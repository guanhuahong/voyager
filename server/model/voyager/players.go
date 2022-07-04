// 自动生成模板Player
package voyager

import "time"

// Player 结构体
// 如果含有time.Time 请自行import time包
type Player struct {
	ID        uint      `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	Username  string    `json:"username" form:"username" gorm:"column:username;comment:用户名;size:64;"`
	Nickname  string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:64;"`
	Mobile    string    `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号码（必要时可作为登录账号或者密码找回）;"`
	Age       *int      `json:"age" form:"age" gorm:"column:age;comment:年龄;"`
	Idno      string    `json:"idno" form:"idno" gorm:"column:idno;comment:身份证号码;"`
}

type PlayerCreated struct {
	Player
	PlayerAuth PlayerAuth `gorm:"foreignkey:PlayerId;association_foreignkey:ID"`
}

// TableName Player 表名
func (Player) TableName() string {
	return "tb_players"
}
