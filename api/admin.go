package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)

//管理员注册接口
func AdminRegisterService(c *gin.Context)  {
	var user service.AdminRegisterService
	if err:=c.ShouldBind(&user);err == nil {
		res := user.AdminRegister()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//管理员登陆接口
func AdminLoginService(c *gin.Context)  {
	var user service.AdminLoginService
	if err:=c.ShouldBind(&user);err==nil{
		res:=user.AdminLogin()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}

}