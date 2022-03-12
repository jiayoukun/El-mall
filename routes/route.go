package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
	"test15/api"
	"test15/middleware"
	"test15/service"
)

//路由配置
func NewRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r:=gin.Default()
	store:=cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession",store))
	r.Use(Recovery)
	//路由
	v1 :=r.Group("/api/v1")
	{
		//用户注册登录
		v1.GET("ws",service.WsHandler)
		v1.POST("user/register",api.RegisterService)
		v1.POST("user/login",api.LoginService)
		//展示所有商品
		v1.GET("products/list",api.Listroduct)
		//展示商品详情
		v1.GET("products/:id",api.ShowProduct)
		v1.GET("electronicrank",api.ShowElectronicRank)
		v1.GET("products/list/1",api.ListCategory1)

		//需要token验证
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//收货地址操作
			authed.POST("address/create",api.CreatAddressService)
			authed.PUT("address/update",api.UpdateAddressService)
			authed.DELETE("address/delete",api.DeleteAddressService)
			//订单操作
			authed.POST("order/create",api.CreatOrderService)
			authed.GET("order/show/:id",api.ShowOrderService)
			authed.PUT("order/update",api.UpdateOrderService)
			authed.DELETE("order/delete",api.DeleteOrderService)
			//购物车操作
			authed.POST("cart/create",api.CreatCartService)
			authed.GET("cart/show/:id",api.ShowCartService)
			authed.PUT("cart/update",api.UpdateCartService)
			authed.DELETE("cart/delete",api.DeleteCartService)
		}
	}
	v2 := r.Group("/api/v2")
	{
		//管理员注册登录
		v2.POST("admin/register",api.AdminRegisterService)
		v2.POST("admin/login",api.AdminLoginService)

		authed2:=v2.Group("/")
		authed2.Use(middleware.JWTAdmin())
		{
		//商品操作
		authed2.POST("products/create/:id",api.CreateProduct)
		authed2.PUT("products/update",api.UpdateProduct)
		authed2.DELETE("products/delete",api.DeleteProduct)

	}
	}
	return r
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logging.Error("gin catch error: ", r)
			c.JSON(http.StatusInternalServerError,"系统内部错误")
		}
	}()
	c.Next()
}
