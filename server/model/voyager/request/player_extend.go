package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PlayerExtendSearch struct{
    voyager.PlayerExtend
    request.PageInfo
    request.SearchDateRangeInfo
}
