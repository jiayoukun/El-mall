package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type DeleteCartService struct {
	Username string `json:"username" form:"username"`
	ProductID string `json:"product_id" form:"product_id"`
}

func (service*DeleteCartService)Delete() serializer.Response {
	var cart model.Cart
	code:=e.SUCCESS
	if err:=model.DB.Model(&cart).Where("username=? AND product_id=?",service.Username,service.ProductID).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_NOT_EXIST_CARTPRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	if err:=model.DB.Model(&cart).Where("username = ? AND product_id = ?",service.Username,service.ProductID).Delete(&cart).Error;err!=nil{
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