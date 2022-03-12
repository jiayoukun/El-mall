package service

import (
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type ListCategory1Service struct {
}

func (server *ListCategory1Service)List() serializer.Response {
	var product []model.Product
	code:=e.SUCCESS
	model.DB.Find(&product,"category_id = ?","1")
	return serializer.Response{
		Status: code,
		Data:serializer.BuildProductData(product),
		Msg: e.GetMsg(code),
	}
}
