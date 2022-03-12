package service

import (
	"fmt"
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type CreateProductService struct {
	Name string `json:"name" form:"name"`
	CategoryId string `json:"category_id" form:"categoryid"`
	Title string `json:"title" form:"title"`
	ImgPath string `json:"img_path" form:"imgpath"`
	Price string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discountprice"`
	Info   string `json:"info" form:"info"`
	Num string `json:"num" form:"num"`
}

func (service*CreateProductService)Create(id string)serializer.Response  {
	var product model.Product
	var count int
	code:=e.SUCCESS
	model.DB.Model(&product).Where("product_id=?",id).Count(&count)
	fmt.Println(count)
	if count !=0 {
		code=e.ERROR_PRODUCT_EXIST
		return serializer.Response{
			Status: code,
			Msg:e.GetMsg(code),
		}
	}

	product=model.Product{
	Name: service.Name,
		CategoryId: service.CategoryId,
		Title: service.Title,
		ImgPath: service.ImgPath,
		Price: service.Price,
		DiscountPrice: service.DiscountPrice,
		Info: service.Info,
		ProductId: id,
		Num: service.Num,
	}
	if err:=model.DB.Create(&product).Error;err!=nil{
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