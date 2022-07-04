// 自动生成模板Player
package voyager

// Player 结构体
// 如果含有time.Time 请自行import time包
type PlayerAuth struct {
	PlayerId *uint  `json:"player_id" form:"player_id" gorm:"column:player_id;comment:玩家ID;"`
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户名;size:64;"`
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
}

// TableName Player 表名
func (PlayerAuth) TableName() string {
	return "tb_player_auths"
}
