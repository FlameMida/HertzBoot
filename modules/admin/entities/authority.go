package entities

import "time"

type Authority struct {
	AuthorityId     string      `json:"authorityId" gorm:"column:authority_id;not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName   string      `json:"authorityName" gorm:"column:authority_name;comment:角色名"`
	ParentId        string      `json:"parentId" gorm:"column:parent_id;comment:父角色ID"`
	DataAuthorityId []Authority `json:"dataAuthorityId" gorm:"many2many:data_authority_id"`
	Children        []Authority `json:"children" gorm:"-"`
	SysBaseMenus    []BaseMenu  `json:"menus" gorm:"many2many:authority_menus;"`
	DefaultRouter   string      `json:"defaultRouter" gorm:"column:default_router;comment:默认菜单;default:dashboard"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}
