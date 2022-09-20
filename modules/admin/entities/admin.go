package entities

import (
	"HertzBoot/app/global"
	"github.com/satori/go.uuid"
)

type Admin struct {
	global.MODEL
	UUID        uuid.UUID   `json:"uuid" gorm:"column:uuid;comment:用户UUID"`
	Username    string      `json:"userName" gorm:"column:username;comment:用户登录名"`
	Password    string      `json:"-"  gorm:"column:password;comment:userName;用户登录密码"`
	NickName    string      `json:"nickName" gorm:"column:nickname;default:系统用户;comment:用户昵称"`
	Avatar      string      `json:"avatar" gorm:"column:avatar;comment:用户头像"`
	Authority   Authority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string      `json:"authorityId" gorm:"column:authority_id;default:888;comment:用户角色ID"`
	SideMode    string      `json:"sideMode" gorm:"column:side_mode;default:dark;comment:边栏模式"`
	ActiveColor string      `json:"activeColor" gorm:"column:active_color;default:#1890ff;comment:用户选择颜色"`
	BaseColor   string      `json:"baseColor" gorm:"column:base_color;default:#fff;comment:基础颜色"` //
	Authorities []Authority `json:"authorities" gorm:"many2many:admin_authority;"`
}
