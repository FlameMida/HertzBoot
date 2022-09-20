package source

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/admin/entities"
	"github.com/gookit/color"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []entities.Admin{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", Avatar: "", AuthorityId: "888"},
}

// Init
// @author:      Flame
// @description: users 表数据初始化
func (a *admin) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]entities.Admin{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> admin 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> admin 表初始数据成功!")
		return nil
	})
}
