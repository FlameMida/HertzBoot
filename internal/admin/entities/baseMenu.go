package entities

import (
	"HertzBoot/pkg/global"
)

type BaseMenu struct {
	global.MODEL
	MenuLevel      uint   `json:"-" gorm:"column:menu_level;comment:菜单等级"`
	ParentId       string `json:"parentId" gorm:"column:parent_id;comment:父菜单ID"`
	Path           string `json:"path" gorm:"column:path;comment:路由path"`
	Name           string `json:"name" gorm:"column:name;comment:路由name"`
	Hidden         bool   `json:"hidden" gorm:"column:hidden;comment:是否在列表隐藏"`
	Component      string `json:"component" gorm:"column:component;comment:对应前端文件路径"`
	Sort           int    `json:"sort" gorm:"column:sort;comment:排序标记"`
	Meta           `json:"meta" gorm:"comment:附加属性"`
	SysAuthorities []Authority         `json:"authorities" gorm:"many2many:authority_menus;"`
	Children       []BaseMenu          `json:"children" gorm:"-"`
	Parameters     []BaseMenuParameter `json:"parameters"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;comment:是否是基础路由"`
	Title       string `json:"title" gorm:"column:title;comment:菜单名"`
	Icon        string `json:"icon" gorm:"column:icon;comment:菜单图标"`
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;comment:自动关闭tab"`
}

type BaseMenuParameter struct {
	global.MODEL
	BaseMenuID uint   `gorm:"column:base_menu_id;comment:基础菜单关联id"`
	Type       string `json:"type" gorm:"column:type;comment:地址栏携带参数为params还是query"`
	Key        string `json:"key" gorm:"column:key;comment:地址栏携带参数的key"`
	Value      string `json:"value" gorm:"column:value;comment:地址栏携带参数的值"`
}
