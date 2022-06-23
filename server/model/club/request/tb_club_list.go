package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TbClubListSearch struct {
	club.TbClubList
	request.PageInfo
	request.SearchDateRangeInfo
}
