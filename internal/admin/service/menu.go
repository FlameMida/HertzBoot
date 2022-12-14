package service

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/request"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

//@author: Flame
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.Menu

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]entities.Menu) {
	var allMenus []entities.Menu
	treeMap = make(map[string][]entities.Menu)
	err = global.DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: Flame
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.Menu

func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []entities.Menu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@author: Flame
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.Menu, treeMap map[string][]model.Menu
//@return: err error

func (menuService *MenuService) getChildrenList(menu *entities.Menu, treeMap map[string][]entities.Menu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: Flame
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []entities.BaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@author: Flame
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.BaseMenu, treeMap map[string][]model.BaseMenu
//@return: err error

func (menuService *MenuService) getBaseChildrenList(menu *entities.BaseMenu, treeMap map[string][]entities.BaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: Flame
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.BaseMenu
//@return: error

func (menuService *MenuService) AddBaseMenu(menu entities.BaseMenu) error {
	if !errors.Is(global.DB.Where("name = ?", menu.Name).First(&entities.BaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.DB.Create(&menu).Error
}

//@author: Flame
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.BaseMenu

func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]entities.BaseMenu) {
	var allMenus []entities.BaseMenu
	treeMap = make(map[string][]entities.BaseMenu)
	err = global.DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: Flame
//@function: GetBaseMenuTree
//@description: 获取基础路由树
//@return: err error, menus []model.BaseMenu

func (menuService *MenuService) GetBaseMenuTree() (err error, menus []entities.BaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@author: Flame
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.BaseMenu, authorityId string
//@return: err error

func (menuService *MenuService) AddMenuAuthority(menus []entities.BaseMenu, authorityId string) (err error) {
	var auth entities.Authority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

//@author: Flame
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.Menu

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []entities.Menu) {
	err = global.DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
