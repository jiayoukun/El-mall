package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)
//展示所有商品接口
func Listroduct(c*gin.Context)  {
	var product service.ListProductService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.List()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//展示单一商品详情接口
func ShowProduct(c*gin.Context)  {
	var product service.ShowProductService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.Show(c.Param("id"))
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//创建商品接口
func CreateProduct(c*gin.Context)  {
	var product service.CreateProductService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.Create(c.Param("id"))
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//更新商品接口
func UpdateProduct(c*gin.Context)  {
	var product service.UpdateProductService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.Update()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//删除商品接口
func DeleteProduct(c*gin.Context)  {
	var product service.DeleteProductService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.Delete()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//展示种类1接口
func ListCategory1(c*gin.Context)  {
	var product service.ListCategory1Service
	if err:=c.ShouldBind(&product);err == nil{
		res := product.List()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
