package entities

import (
	"HertzBoot/app/global"
)

type JwtBlacklist struct {
	global.MODEL
	Jwt string `gorm:"column:jwt;type:text;comment:jwt"`
}
