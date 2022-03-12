package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type CreateAddressService struct {
	Username string `json:"username" form:"username"`
	Phone string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	Name string `json:"name" form:"name"`
}

func (service*CreateAddressService)Create()serializer.Response  {
	var address model.Address
	//var count int
	code:=e.SUCCESS

	address.Username=service.Username
	address.Phone=service.Phone
	address.Name=service.Name
	address.Address=service.Address
	if err:=model.DB.Create(&address).Error;err!=nil{
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