package responses

import (
	"HertzBoot/internal/admin/entities"
)

type SysMenusResponse struct {
	Menus []entities.Menu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []entities.BaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu entities.BaseMenu `json:"menu"`
}
