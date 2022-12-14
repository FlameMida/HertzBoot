package source

import (
	"HertzBoot/pkg/global"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var DataAuthorities = new(dataAuthorities)

type dataAuthorities struct{}

type DataAuthority struct {
	AuthorityId   string `gorm:"column:authority_authority_id"`
	DataAuthority string `gorm:"column:data_authority_id_authority_id"`
}

var infos = []DataAuthority{
	{"888", "888"},
}

// Init @author: Flame
// @description: data_authority_id 表数据初始化
func (d *dataAuthorities) Init() error {
	return global.DB.Table("data_authority_id").Transaction(func(tx *gorm.DB) error {
		if tx.Where("data_authority_id_authority_id IN ('888') ").Find(&[]DataAuthority{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> data_authority_id 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> data_authority_id 表初始数据成功!")
		return nil
	})
}
