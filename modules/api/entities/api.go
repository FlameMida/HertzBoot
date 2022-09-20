package entities

import (
	"HertzBoot/app/global"
)

type Api struct {
	global.MODEL
	Path        string `json:"path" gorm:"column:path;comment:api路径"`
	Description string `json:"description" gorm:"column:description;comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"column:api_group;comment:api组"`
	Method      string `json:"method" gorm:"column:method;default:POST;comment:方法"`
}
