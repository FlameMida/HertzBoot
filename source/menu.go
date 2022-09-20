package source

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/admin/entities"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []entities.BaseMenu{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: entities.Meta{Title: "仪表盘", Icon: "setting"}},
	{MODEL: global.MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: entities.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{MODEL: global.MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: entities.Meta{Title: "角色管理", Icon: "s-custom"}},
	{MODEL: global.MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: entities.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{MODEL: global.MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: entities.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{MODEL: global.MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: entities.Meta{Title: "用户管理", Icon: "coordinate"}},

	{MODEL: global.MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: entities.Meta{Title: "系统工具", Icon: "s-cooperation"}},

	{MODEL: global.MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "8", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: entities.Meta{Title: "系统配置", Icon: "s-operation"}},

	{MODEL: global.MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/operations.vue", Sort: 6, Meta: entities.Meta{Title: "操作历史", Icon: "time"}},

	{MODEL: global.MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/entities/state.vue", Sort: 6, Meta: entities.Meta{Title: "服务器状态", Icon: "cloudy"}},
}

// Init @author: Flame
// @description: base_menus 表数据初始化
func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]entities.BaseMenu{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> base_menus 表初始数据成功!")
		return nil
	})
}
