package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
)

type PlayerSearch struct {
	voyager.Player
	request.PageInfo
	request.SearchDateRangeInfo
}

type PlayerForm struct {
	voyager.Player
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
}
