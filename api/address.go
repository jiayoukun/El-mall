package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)

//创建收货地址
func CreatAddressService(c*gin.Context)  {
	var address service.CreateAddressService
	if err:=c.ShouldBind(&address);err==nil{
		res:=address.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//修改收货地址
func UpdateAddressService(c*gin.Context)  {
	var address service.UpdateAddressService
	if err:=c.ShouldBind(&address);err==nil{
		res:=address.Update()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//删除收货地址
func DeleteAddressService(c*gin.Context)  {
	var address service.DeleteAddressService
	if err:=c.ShouldBind(&address);err==nil{
		res:=address.Delete()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}