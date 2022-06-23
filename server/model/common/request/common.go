package request

import "time"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

type SearchDateRangeInfo struct {
	CreatedAtStart time.Time `json:"dtcs" form:"dtcs" time_format:"2006-01-02 15:04:05"`
	CreatedAtEnd   time.Time `json:"dtce" form:"dtce" time_format:"2006-01-02 15:04:05"`
	UpdatedAtStart time.Time `json:"dtus" form:"dtus" time_format:"2006-01-02 15:04:05"`
	UpdatedAtEnd   time.Time `json:"dtue" form:"dtue" time_format:"2006-01-02 15:04:05"`
}
