package service

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/pkg/global"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type BaseMenuService struct {
}

// @author:      Flame
// @function:    DeleteBaseMenu
// @description: 删除基础路由
// @param:       id float64
// @return:      err error

func (baseMenuService *BaseMenuService) DeleteBaseMenu(id float64) (err error) {
	err = global.DB.Preload("Parameters").Where("parent_id = ?", id).First(&entities.BaseMenu{}).Error
	if err != nil {
		var menu entities.BaseMenu
		db := global.DB.Preload("SysAuthorities").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.DB.Delete(&entities.BaseMenuParameter{}, "base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthorities) > 0 {
			err = global.DB.Model(&menu).Association("SysAuthorities").Delete(&menu.SysAuthorities)
			if err != nil {
				return err
			}
		} else {
			err = db.Error
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// @author:      Flame
// @function:    UpdateBaseMenu
// @description: 更新路由
// @param:       menu model.BaseMenu
// @return:      err error

func (baseMenuService *BaseMenuService) UpdateBaseMenu(menu entities.BaseMenu) (err error) {
	var oldMenu entities.BaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&entities.BaseMenu{}).Error, gorm.ErrRecordNotFound) {
				hlog.Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&entities.BaseMenuParameter{}, "base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			hlog.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].BaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				hlog.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			hlog.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}

// @author:      Flame
// @function:    GetBaseMenuById
// @description: 返回当前选中menu
// @param:       id float64
// @return:      err error, menu model.BaseMenu

func (baseMenuService *BaseMenuService) GetBaseMenuById(id float64) (err error, menu entities.BaseMenu) {
	err = global.DB.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
