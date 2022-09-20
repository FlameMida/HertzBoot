package responses

import (
	"HertzBoot/modules/core/http/requests"
)

type PolicyPathResponse struct {
	Paths []requests.CasbinInfo `json:"paths"`
}
