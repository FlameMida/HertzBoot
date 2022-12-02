package service

import (
	"HertzBoot/config"
	"HertzBoot/internal/core/entities"
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

//@author: Flame
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type ConfigService struct {
}

func (configService *ConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.CONFIG
}

// @description set system config,
//@author: Flame
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (configService *ConfigService) SetSystemConfig(system entities.System) (err error) {
	cs := pkg.StructToMap(system.Config)
	for k, v := range cs {
		global.VP.Set(k, v)
	}
	err = global.VP.WriteConfig()
	return err
}

//@author: Flame
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (configService *ConfigService) GetServerInfo() (server *pkg.Server, err error) {
	var s pkg.Server
	s.Os = pkg.InitOS()
	if s.Cpu, err = pkg.InitCPU(); err != nil {
		hlog.Error("func utils.InitCPU() Failed", err.Error())
		return &s, err
	}
	if s.Rrm, err = pkg.InitRAM(); err != nil {
		hlog.Error("func utils.InitRAM() Failed", err.Error())
		return &s, err
	}
	if s.Disk, err = pkg.InitDisk(); err != nil {
		hlog.Error("func utils.InitDisk() Failed", err.Error())
		return &s, err
	}

	return &s, nil
}
