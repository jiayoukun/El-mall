package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type ShowOrderService struct {
}

func (service*ShowOrderService)Show(id string) serializer.Response {
	var order []model.Order
	var count int
	code:=e.SUCCESS
	if err:=model.DB.Model(&order).Where("username=?",id).Count(&count).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if count==0 {
		code=e.ERROR_NOT_EXIST_ORDER
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}else {
		model.DB.Find(&order,"username=?",id)
	}
	return serializer.Response{
		Status: code,
		Data: serializer.BuildOrderData(order),
		Msg: e.GetMsg(code),
	}
}