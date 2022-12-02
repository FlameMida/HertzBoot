package responses

import "HertzBoot/internal/admin/entities"

type SysAuthorityResponse struct {
	Authority entities.Authority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      entities.Authority `json:"authority"`
	OldAuthorityId string             `json:"oldAuthorityId"` // 旧角色ID
}
