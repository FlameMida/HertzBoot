package source

import (
	"HertzBoot/pkg/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId string `gorm:"column:authority_authority_id"`
	BaseMenuId  uint   `gorm:"column:base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{"888", 1},
	{"888", 3},
	{"888", 4},
	{"888", 5},
	{"888", 6},
	{"888", 7},
	{"888", 8},
	{"888", 9},
	{"888", 10},
	{"888", 11},
}

// Init @author: Flame
// @description: authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.DB.Table("authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_authority_id IN ('888', '8881', '9528')").Find(&[]AuthorityMenus{}).RowsAffected == 48 {
			color.Danger.Println("\n[Mysql] --> authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> authority_menus 表初始数据成功!")
		return nil
	})
}
