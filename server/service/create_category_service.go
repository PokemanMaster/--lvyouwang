package service

import (
	"qiqi-go/models"
	"qiqi-go/pkg/e"
	"qiqi-go/pkg/logging"
	"qiqi-go/serializer"
)

// CreateCategoryService 收藏创建的服务
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// Create 创建分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := models.Category{
		CategoryID:   service.CategoryID,
		CategoryName: service.CategoryName,
	}
	code := e.SUCCESS
	err := models.DB.Create(&category).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCategory(category),
	}
}
