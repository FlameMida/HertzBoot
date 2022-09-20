package service

import (
	"HertzBoot/app/global"
	"HertzBoot/app/request"
	"HertzBoot/modules/admin/entities"
	"HertzBoot/tools"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// @author:      Flame
// @function:    Register
// @description: 用户注册
// @param:       u model.Admin
// @return:      err error, userInter model.Admin

type AdminService struct{}

func (adminService *AdminService) Register(u entities.Admin) (err error, userInter entities.Admin) {
	var user entities.Admin
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = tools.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return err, u
}

// @author:      Flame
// @function:    ChangePassword
// @description: 修改用户密码
// @param:       u *model.Admin, newPassword string
// @return:      err error, userInter *model.Admin

func (adminService *AdminService) ChangePassword(u *entities.Admin, newPassword string) (err error, userInter *entities.Admin) {
	var user entities.Admin
	u.Password = tools.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", tools.MD5V([]byte(newPassword))).Error
	return err, u
}

// @author:      Flame
// @function:    GetUserInfoList
// @description: 分页获取数据
// @param:       info request.PageInfo
// @return:      err error, list interface{}, total int64

func (adminService *AdminService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&entities.Admin{})
	var userList []entities.Admin
	err = db.Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @author:      Flame
// @function:    SetUserAuthority
// @description: 设置一个用户的权限
// @param:       uuid uuid.UUID, authorityId string
// @return:      err error

func (adminService *AdminService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.DB.Where("user_id = ? AND authority_authority_id = ?", id, authorityId).First(&entities.AdminAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.DB.Where("uuid = ?", uuid).First(&entities.Admin{}).Update("authority_id", authorityId).Error
	return err
}

// @author:      Flame
// @function:    SetUserAuthorities
// @description: 设置一个用户的权限
// @param:       id uint, authorityIds []string
// @return:      err error

func (adminService *AdminService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]entities.AdminAuthority{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []entities.AdminAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, entities.AdminAuthority{
				AdminId: id, AuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

// @author:      Flame
// @function:    DeleteUser
// @description: 删除用户
// @param:       id float64
// @return:      err error

func (adminService *AdminService) DeleteUser(id float64) (err error) {
	var user entities.Admin
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(&[]entities.AdminAuthority{}, "user_id = ?", id).Error
	return err
}

// @author:      Flame
// @function:    SetUserInfo
// @description: 设置用户信息
// @param:       reqUser model.Admin
// @return:      err error, user model.Admin

func (adminService *AdminService) SetUserInfo(reqUser entities.Admin) (err error, user entities.Admin) {
	err = global.DB.Updates(&reqUser).Error
	return err, reqUser
}

// @author:      Flame
// @function:    GetUserInfo
// @description: 获取用户信息
// @param:       uuid uuid.UUID
// @return:      err error, user entities.Admin

func (adminService *AdminService) GetUserInfo(uuid uuid.UUID) (err error, user entities.Admin) {
	var reqUser entities.Admin
	err = global.DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

// @author:      Flame
// @function:    FindUserById
// @description: 通过id获取用户信息
// @param:       id int
// @return:      err error, user *model.Admin

func (adminService *AdminService) FindUserById(id int) (err error, user *entities.Admin) {
	var u entities.Admin
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// @author:      Flame
// @function:    FindUserByUuid
// @description: 通过uuid获取用户信息
// @param:       uuid string
// @return:      err error, user *model.Admin

func (adminService *AdminService) FindUserByUuid(uuid string) (err error, user *entities.Admin) {
	var u entities.Admin
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
