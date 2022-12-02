package entities

import (
	"HertzBoot/pkg/global"
)

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"column:jwt;type:text;comment:jwt"`
}
