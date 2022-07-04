// 自动生成模板PlayerExtend
package voyager

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PlayerExtend 结构体
// 如果含有time.Time 请自行import time包
type PlayerExtend struct {
      global.GVA_MODEL
      PlayerId  *int `json:"playerId" form:"playerId" gorm:"column:player_id;comment:玩家id;size:19;"`
      Field  string `json:"field" form:"field" gorm:"column:field;comment:扩展字段名;size:64;"`
      Value  string `json:"value" form:"value" gorm:"column:value;comment:扩展字段值;size:255;"`
}


// TableName PlayerExtend 表名
func (PlayerExtend) TableName() string {
  return "tb_player_extend"
}

