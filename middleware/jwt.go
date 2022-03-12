package middleware

import (
	"github.com/gin-gonic/gin"
	"test15/pkg/e"
	"test15/pkg/utils"
	"time"
)

//验证token中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code:=200
		var data interface{}
		token:=c.GetHeader("Authorization")
		if token==""{
			code=404
		}else{
			claims,err:=utils.Checktoken(token)
			if err != nil {
				code=e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix()>claims.ExpiresAt{
				code=e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}else if claims.Authority ==1 {
				code = 404
			}
		}
		if code != e.SUCCESS {
			c.JSON(200,gin.H{
				"status":200,
				"msg":"token验证失败",
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
//验证管理员token中间件
func JWTAdmin()gin.HandlerFunc  {
	return func(c *gin.Context) {
		code:=200
		var data interface{}
		token:=c.GetHeader("Authorization")
		if token==""{
			code=e.INVALID_PARAMS
		}else{
			claims,err:=utils.Checktoken(token)
			if err != nil {
				code=e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix()>claims.ExpiresAt{
				code=e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}else if claims.Authority ==0 {
				code = e.ERROR_AUTH_INSUFFICIENT_AUTHORITY
			}
		}
		if code != e.SUCCESS {
			c.JSON(200,gin.H{
				"status":200,
				"msg":"token验证失败",
				"data":data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}