package requests

import (
	"HertzBoot/app/request"
	"HertzBoot/modules/api/entities"
)

// SearchApiParams api分页条件查询及排序结构体
type SearchApiParams struct {
	entities.Api
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
