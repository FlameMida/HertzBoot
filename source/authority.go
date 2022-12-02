package source

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/pkg/global"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

var authorities = []entities.Authority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "超级管理员", ParentId: "0", DefaultRouter: "dashboard"},
}

// Init @author: Flame
// @description: authorities 表数据初始化
func (a *authority) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888"}).Find(&[]entities.Authority{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> authorities 表初始数据成功!")
		return nil
	})
}
