package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)

//加入购物车接口
func CreatCartService(c *gin.Context)  {
	var cart service.CreateCartService
	if err:=c.ShouldBind(&cart);err == nil {
		res:=cart.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//修改购物车接口
func UpdateCartService(c *gin.Context)  {
	var cart service.UpdateCartService
	if err:=c.ShouldBind(&cart);err == nil {
		res:=cart.Update()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//删除购物车接口
func DeleteCartService(c *gin.Context)  {
	var cart service.DeleteCartService
	if err:=c.ShouldBind(&cart);err == nil {
		res:=cart.Delete()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
//展示购物车接口
func ShowCartService(c *gin.Context)  {
	var cart service.ShowCartService
	if err:=c.ShouldBind(&cart);err == nil {
		res:=cart.Show(c.Param("id"))
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}