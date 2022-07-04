package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/voyager"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PlayerProfileSearch struct{
    voyager.PlayerProfile
    request.PageInfo
    request.SearchDateRangeInfo
}
