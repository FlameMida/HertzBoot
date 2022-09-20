package service

import (
	"HertzBoot/app/global"
	"HertzBoot/app/request"
	"HertzBoot/modules/api/entities"
	"HertzBoot/modules/core/service"
	"errors"

	"gorm.io/gorm"
)

//@author: Flame
//@function: CreateApi
//@description: 新增基础api
//@param: api model.Api
//@return: err error

type ApiService struct {
}

var casbinService = new(service.CasbinService)

func (apiService *ApiService) CreateApi(api entities.Api) (err error) {
	if !errors.Is(global.DB.
		Where("path = ? AND method = ?", api.Path, api.Method).
		First(&entities.Api{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.DB.Create(&api).Error
}

//@author: Flame
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.Api
//@return: err error

func (apiService *ApiService) DeleteApi(api entities.Api) (err error) {
	err = global.DB.Delete(&api).Error
	casbinService.ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: Flame
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.Api, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(api entities.Api, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&entities.Api{})
	var apiList []entities.Api

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: Flame
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.Api

func (apiService *ApiService) GetAllApis() (err error, apis []entities.Api) {
	err = global.DB.Find(&apis).Error
	return
}

//@author: Flame
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.Api

func (apiService *ApiService) GetApiById(id float64) (err error, api entities.Api) {
	err = global.DB.Where("id = ?", id).First(&api).Error
	return
}

//@author: Flame
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.Api
//@return: err error

func (apiService *ApiService) UpdateApi(api entities.Api) (err error) {
	var oldA entities.Api
	err = global.DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&entities.Api{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = casbinService.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.DB.Save(&api).Error
		}
	}
	return err
}

//@author: Flame
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.Api
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]entities.Api{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return global.DB.Delete(entities.Api{}, ids).Error
}
