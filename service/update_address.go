package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type UpdateAddressService struct {
	Username string `json:"username" form:"username"`
	Phone string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	Name string `json:"name" form:"name"`
}

func (service*UpdateAddressService)Update() serializer.Response {
	var address model.Address
	code:=e.SUCCESS
	var count int
	if err:=model.DB.First(&address,"username=?",service.Username).Count(&count).Error;err!=nil{
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
	if service.Phone!=""{
		address.Phone=service.Phone
	}
	if service.Address!=""{
		address.Address=service.Address
	}
	if service.Name!=""{
		address.Name=service.Name
	}
	if err:=model.DB.Model(&address).Update(&address).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error:err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}