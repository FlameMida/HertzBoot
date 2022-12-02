package responses

import "HertzBoot/internal/api/entities"

type SysAPIResponse struct {
	Api entities.Api `json:"api"`
}

type SysAPIListResponse struct {
	Apis []entities.Api `json:"apis"`
}
