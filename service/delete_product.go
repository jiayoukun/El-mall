package service

import (
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type DeleteProductService struct {
	ProductID string `json:"product_id" form:"product_id"`
}

func (service*DeleteProductService)Delete()serializer.Response  {
	var product model.Product
	code:=e.SUCCESS
	var count int
	model.DB.Model(&product).Where("product_id=?",service.ProductID).Count(&count)
	if count == 0 {
		code=e.ERROR_NOT_EXIST_PRODUCT
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	err := model.DB.Where("product_id=?",service.ProductID).Delete(&product).Error
	if err!= nil{
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