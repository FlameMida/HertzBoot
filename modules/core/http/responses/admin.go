package responses

import (
	"HertzBoot/modules/admin/entities"
)

type UserResponse struct {
	User entities.Admin `json:"user"`
}

type LoginResponse struct {
	User      entities.Admin `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
