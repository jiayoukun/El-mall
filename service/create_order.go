package service

import (
	"fmt"
	logging "github.com/sirupsen/logrus"
	"math/rand"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
	"time"
)

type CreateOrderService struct {
	Username string `json:"user_id" form:"username"`
	ProductID string `json:"product_id" form:"product_id"`
	Num string `json:"num" form:"num"`
	Sumprice string `json:"sumprice" form:"sumprice"`
}

func (service*CreateOrderService)Create()serializer.Response  {
	var order model.Order
	var address model.Address
	var count int
	code:=e.SUCCESS
	model.DB.Model(&address).Where("username=?",service.Username).Count(&count)
	if count == 0 {
		code=e.ERROR_NOT_EXIST_ADDRESS
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	model.DB.First(&address,"username=?",service.Username)
	order.Address=address.Address
	order.AddressName=address.Name
	order.AddressPhone=address.Phone
	order.Username=service.Username
	order.ProductID=service.ProductID
	order.Num=service.Num
	order.SumPrice=service.Sumprice
	//生成随机订单号
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	number=number+service.ProductID+service.Username
	order.OrderNum=number
	if err:=model.DB.Create(&order).Error;err!=nil{
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