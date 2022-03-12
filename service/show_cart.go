package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type ShowCartService struct {
}

func (service*ShowCartService)Show(id string) serializer.Response {
	var cart []model.Cart
	var count int
	code:=e.SUCCESS
	if err:=model.DB.Model(&cart).Where("username=?",id).Count(&count).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if count==0 {
		code=e.ERROR_NOT_EXIST_CARTPRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}else {
		model.DB.Find(&cart,"username=?",id)
	}
	return serializer.Response{
		Status: code,
		Data: serializer.BuildCartData(cart),
		Msg: e.GetMsg(code),
	}
}