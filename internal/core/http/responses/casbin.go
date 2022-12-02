package responses

import (
	"HertzBoot/internal/core/http/requests"
)

type PolicyPathResponse struct {
	Paths []requests.CasbinInfo `json:"paths"`
}
