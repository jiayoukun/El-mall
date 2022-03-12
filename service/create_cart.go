package service

import (
	logging "github.com/sirupsen/logrus"
	"strconv"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type CreateCartService struct {
	Username string `json:"user_id" form:"username"`
	ProductID string `json:"product_id" form:"product_id"`
	Num string `json:"num" form:"num"`
}

func (service*CreateCartService)Create() serializer.Response {
	var cart model.Cart
	code:=e.SUCCESS
	var count int
	if err:=model.DB.Model(&cart).Where("username=? AND product_id=?",service.Username,service.ProductID).Count(&count).Error;err!=nil{
		logging.Info(err)
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	num,_:=strconv.ParseInt(service.Num,10,64)
	cartnum,_:=strconv.ParseInt(cart.Num,10,64)
	maxnum,_:=strconv.ParseInt(cart.MaxNum,10,64)
	if num>maxnum {
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	var product model.Product
	if err:=model.DB.First(&product,"product_id=?",service.ProductID).Error;err!=nil{
		code=e.ERROR_DATABASE
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	cart.ProductID=product.ProductId
	cart.Username=service.Username
	if service.Num==""{
		cart.Num="1"
	}else {
		cart.Num=service.Num
	}
	cart.MaxNum=product.Num
	cart.ProductName=product.Name
	cart.Imgpath=product.ImgPath
	cart.Price=product.DiscountPrice
	if count==0 {
		model.DB.Model(&cart).Create(&cart)
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}else if cartnum+num<maxnum {
		model.DB.Model(&cart).Update(&cart)
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
	}
}