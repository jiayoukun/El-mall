package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type DeleteOrderService struct {
	Username string `json:"username" form:"username"`
	Ordernum string `json:"order_num" form:"order_num"`
}

func (service*DeleteOrderService)Delete() serializer.Response {
	var order model.Order
	code:=e.SUCCESS
	var count int
	if err:=model.DB.Find(&order,"username = ? AND order_num = ?",service.Username,service.Ordernum).Count(&count).Error;err!=nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err:=model.DB.Model(&order).Where("username = ? AND order_num = ?",service.Username,service.Ordernum).Delete(&order).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}