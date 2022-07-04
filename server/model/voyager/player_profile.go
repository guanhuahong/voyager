// 自动生成模板PlayerProfile
package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PlayerProfile 结构体
// 如果含有time.Time 请自行import time包
type PlayerProfile struct {
      global.GVA_MODEL
      PlayerId  *int `json:"playerId" form:"playerId" gorm:"column:player_id;comment:玩家id;size:19;"`
      Nick  string `json:"nick" form:"nick" gorm:"column:nick;comment:玩家昵称;size:64;"`
      Avatar  string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`
      Gold  *float64 `json:"gold" form:"gold" gorm:"column:gold;comment:玩家账户金额;size:19;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:玩家状态;"`
      TotalWinlost  *float64 `json:"totalWinlost" form:"totalWinlost" gorm:"column:total_winlost;comment:累计输赢;size:19;"`
      TotalSysfee  *float64 `json:"totalSysfee" form:"totalSysfee" gorm:"column:total_sysfee;comment:累计抽水;size:19;"`
      TotalBet  *float64 `json:"totalBet" form:"totalBet" gorm:"column:total_bet;comment:累计投注;size:19;"`
      TotalEffbet  *float64 `json:"totalEffbet" form:"totalEffbet" gorm:"column:total_effbet;comment:累计有效投注;size:19;"`
}


// TableName PlayerProfile 表名
func (PlayerProfile) TableName() string {
  return "tb_player_profile"
}

