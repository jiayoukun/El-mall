package service

import (
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type ListProductService struct {
}

func (service*ListProductService)List() serializer.Response {
	var product []model.Product
	code:=e.SUCCESS
	model.DB.Find(&product)
	return serializer.Response{
		Status: code,
		Data:serializer.BuildProductData(product),
		Msg: e.GetMsg(code),
	}
}