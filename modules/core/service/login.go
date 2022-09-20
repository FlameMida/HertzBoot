package service

import (
	"HertzBoot/app/global"
	"HertzBoot/modules/admin/entities"
	"HertzBoot/tools"
)

type LoginService struct {
}

// Login
// @author:      Flame
// @function:    Login
// @description: 后台用户登录
// @param:       u *model.User
// @return:      err error, userInter *model.User
func (loginService *LoginService) Login(u *entities.Admin) (err error, userInter *entities.Admin) {
	var user entities.Admin
	u.Password = tools.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}
