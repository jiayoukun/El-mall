package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"test15/service"
)

//展示单一商品详情接口
func ShowElectronicRank(c*gin.Context)  {
	var product service.ShowElectronicRankService
	if err:=c.ShouldBind(&product);err == nil{
		res := product.Show()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
		logging.Info(err)
	}
}
