package responses

import "HertzBoot/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
