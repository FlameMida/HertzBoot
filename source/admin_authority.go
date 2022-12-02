package source

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/pkg/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AdminAuthority = new(adminAuthority)

type adminAuthority struct{}

var adminAuthorityModel = []entities.AdminAuthority{
	{AdminId: 1, AuthorityAuthorityId: "888"},
}

// Init @description: admin_authority 数据初始化
func (a *adminAuthority) Init() error {
	return global.DB.Model(&entities.AdminAuthority{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("admin_id IN (1)").Find(&[]entities.AdminAuthority{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> admin_authority 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&adminAuthorityModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> admin_authority 表初始数据成功!")
		return nil
	})
}
