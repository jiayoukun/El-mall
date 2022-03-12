package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type DeleteAddressService struct {
	Username string `json:"username" form:"username"`
}

func (service*DeleteAddressService)Delete() serializer.Response {
	var addres model.Address
	code:=e.SUCCESS
	var count int
	if err:=model.DB.First(&addres,"username=?",service.Username).Count(&count).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if count==0 {
		code=e.ERROR_NOT_EXIST_ADDRESS
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	if err:=model.DB.Model(&addres).Delete(&addres).Error;err!=nil{
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