package service

import (
	"HertzBoot/internal/admin/entities"
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
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
	u.Password = pkg.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}
