package service

import (
	"HertzBoot/internal/core/entities"
	"HertzBoot/internal/core/http/requests"
	"HertzBoot/pkg/global"
	"HertzBoot/pkg/request"
)

// @author:      Flame
// @function:    CreateOperations
// @description: 创建记录
// @param:       operations model.Operations
// @return:      err error

type OperationsService struct {
}

func (OperationsService *OperationsService) CreateOperations(operations entities.Operations) (err error) {
	err = global.DB.Create(&operations).Error
	return err
}

// @author:      Flame
// @author:      Flame
// @function:    DeleteOperationsByIds
// @description: 批量删除记录
// @param:       ids request.IdsReq
// @return:      err error

func (OperationsService *OperationsService) DeleteOperationsByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]entities.Operations{}, "id in (?)", ids.Ids).Error
	return err
}

// @author:      Flame
// @function:    DeleteOperations
// @description: 删除操作记录
// @param:       operations model.Operations
// @return:      err error

func (OperationsService *OperationsService) DeleteOperations(operations entities.Operations) (err error) {
	err = global.DB.Delete(&operations).Error
	return err
}

// @author:      Flame
// @function:    DeleteOperations
// @description: 根据id获取单条操作记录
// @param:       id uint
// @return:      err error, operations model.Operations

func (OperationsService *OperationsService) GetOperations(id uint) (err error, operations entities.Operations) {
	err = global.DB.Where("id = ?", id).First(&operations).Error
	return
}

// @author:      Flame
// @author:      Flame
// @function:    GetOperationsInfoList
// @description: 分页获取操作记录列表
// @param:       info requests.OperationsSearch
// @return:      err error, list interface{}, total int64

func (OperationsService *OperationsService) GetOperationsInfoList(info requests.OperationsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&entities.Operations{})
	var operations []entities.Operations
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&operations).Error
	return err, operations, total
}
