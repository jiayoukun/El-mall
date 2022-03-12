package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)

//创建订单接口
func CreatOrderService(c *gin.Context)  {
	var order service.CreateOrderService
	if err:=c.ShouldBind(&order);err == nil {
		res:=order.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//修改订单接口
func UpdateOrderService(c *gin.Context)  {
	var order service.UpdateOrderService
	if err:=c.ShouldBind(&order);err == nil {
		res:=order.Update()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//删除订单接口
func DeleteOrderService(c *gin.Context)  {
	var order service.DeleteOrderService
	if err:=c.ShouldBind(&order);err == nil {
		res:=order.Delete()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//展示我的订单接口
func ShowOrderService(c *gin.Context)  {
	var order service.ShowOrderService
	if err:=c.ShouldBind(&order);err == nil {
		res:=order.Show(c.Param("id"))
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}