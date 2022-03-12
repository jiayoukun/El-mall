package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type UpdateProductService struct {
	Name string `json:"name" form:"name"`
	CategoryId string `json:"category_id" form:"category_id"`
	Title string `json:"title" form:"title"`
	ImgPath string `json:"img_path" form:"imgpath"`
	Price string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discountprice"`
	Info   string `json:"info" form:"info"`
	ProductID string `json:"product_id" form:"product_id"`
	Num string `json:"num" form:"num"`
}

func (service*UpdateProductService)Update()serializer.Response {
	var product model.Product
	var count int
	code:=e.SUCCESS

	model.DB.Model(&product).Where("product_id=?",service.ProductID).Count(&count)
	if count==0{
		code=e.ERROR_NOT_EXIST_PRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	if service.Name != "" {
		product.Name=service.Name
	}
	if service.CategoryId != "" {
		product.CategoryId=service.CategoryId
	}
	if service.Title != "" {
		product.Title=service.Title
	}
	if service.ImgPath != "" {
		product.ImgPath=service.ImgPath
	}
	if service.Price != "" {
		product.Price=service.Price
	}
	if service.DiscountPrice != "" {
		product.DiscountPrice=service.DiscountPrice
	}
	if service.Info != "" {
		product.Info=service.Info
	}
	if service.ProductID != "" {
		product.ProductId=service.ProductID
	}
	if service.Num != "" {
		product.Num=service.Num
	}

	if err:=model.DB.Model(&product).Where("product_id=?",service.ProductID).Update(&product).Error;err!=nil{
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