package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/club"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TbClubMemSearch struct {
	club.TbClubMem
	request.PageInfo
	request.SearchDateRangeInfo
}
