package requests

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/admin/entities"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []entities.BaseMenu
	AuthorityId string // 角色ID
}

func DefaultMenu() []entities.BaseMenu {
	return []entities.BaseMenu{{
		MODEL:     global.MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: entities.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
