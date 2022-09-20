package service

import (
	"HertzBoot/app/global"
	"HertzBoot/config"
	"HertzBoot/modules/core/entities"
	"HertzBoot/tools"
	"go.uber.org/zap"
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
	cs := tools.StructToMap(system.Config)
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

func (configService *ConfigService) GetServerInfo() (server *tools.Server, err error) {
	var s tools.Server
	s.Os = tools.InitOS()
	if s.Cpu, err = tools.InitCPU(); err != nil {
		global.LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = tools.InitRAM(); err != nil {
		global.LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = tools.InitDisk(); err != nil {
		global.LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
