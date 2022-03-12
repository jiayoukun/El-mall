package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type UpdateOrderService struct {
	Username string `json:"user_id" form:"username"`
	Ordernum string `json:"ordernum" form:"order_num"`
	Address string `json:"address" form:"addres"`
	Addressname string `json:"addressname" form:"addressname"`
	Addressphone string `json:"addressphone" form:"addressphone"`
	Num string `json:"num" form:"num"`
	Sumprice string `json:"sumprice" form:"sumprice"`
}

func (service*UpdateOrderService)Update() serializer.Response {
	var order model.Order
	var address model.Address
	var count int
	code:=e.SUCCESS
	if err:=model.DB.First(&order,"order_num=? AND username=?",service.Ordernum,service.Username).Count(&count).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: count,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if count==0 {
		code=e.ERROR_NOT_EXIST_ORDER
		return serializer.Response{
			Status: count,
			Msg: e.GetMsg(code),
		}
	}
	if err:=model.DB.First(&address,"username=?",service.Username).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	if service.Address!="" {
		order.Address=service.Address
	}
	if service.Addressname != "" {
		order.AddressName=service.Addressname
	}
	if service.Addressphone != "" {
		order.AddressPhone=service.Addressphone
	}
	if service.Num!=""{
		order.Num=service.Num
	}
	if service.Sumprice!=""{
		order.SumPrice=service.Sumprice
	}
	if err:=model.DB.Model(&order).Update(&order).Error;err!=nil{
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