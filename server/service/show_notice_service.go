package service

import (
	"qiqi-go/models"
	"qiqi-go/pkg/e"
	"qiqi-go/pkg/logging"
	"qiqi-go/serializer"
)

// ShowNoticeService 公告详情服务
type ShowNoticeService struct {
}

// Show 公告详情服务
func (service *ShowNoticeService) Show() serializer.Response {
	var notice models.Notice
	code := e.SUCCESS
	if err := models.DB.First(&notice, 1).Error; err != nil {
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
		Data:   serializer.BuildNotice(notice),
	}
}
