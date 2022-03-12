package service

import (
	logging "github.com/sirupsen/logrus"
	"strconv"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type UpdateCartService struct {
	Username string `json:"user_id" form:"username"`
	ProductID string `json:"product_id" form:"product_id"`
	Num string `json:"num" form:"num"`
}

func (service*UpdateCartService)Update() serializer.Response {
	var cart model.Cart
	code:=e.SUCCESS
	if err:=model.DB.Where("username=? AND product_id=?",service.Username,service.ProductID).Find(&cart).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_NOT_EXIST_CARTPRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	num,_:=strconv.ParseInt(service.Num,10,64)
	maxnum,_:=strconv.ParseInt(cart.MaxNum,10,64)
	if service.Num!="" && num<=maxnum {
		cart.Num=service.Num
	}else if num>maxnum{
		code=e.ERROR_NOT_ENOUGH_PRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	if err:=model.DB.Model(&cart).Where("username=? AND product_id=?",service.Username,service.ProductID).Update(&cart).Error;err!=nil{
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