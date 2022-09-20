package entities

type Menu struct {
	BaseMenu
	MenuId      string              `json:"menuId" gorm:"column:menu_id;comment:菜单ID"`
	AuthorityId string              `json:"-" gorm:"column:authority_id;comment:角色ID"`
	Children    []Menu              `json:"children" gorm:"-"`
	Parameters  []BaseMenuParameter `json:"parameters" gorm:"foreignKey:BaseMenuID;references:MenuId"`
}

func (s Menu) TableName() string {
	return "authority_menu"
}
