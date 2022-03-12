package api

import (
	"github.com/gin-gonic/gin"
	"test15/service"

	logging "github.com/sirupsen/logrus"
)

//用户注册接口
func RegisterService(c *gin.Context)  {
	var user service.RegisterService
	if err:=c.ShouldBind(&user);err == nil {
		res := user.Register()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}

//用户登陆接口
func LoginService(c *gin.Context)  {
	var user service.LoginService
	if err:=c.ShouldBind(&user);err==nil{
		res:=user.Login()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}

}
