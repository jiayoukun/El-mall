package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
	"test15/model"
	"test15/pkg/utils"
	"test15/serializer"
)

type AdminRegisterService struct {
	Username string	`json:"username" form:"username" binding:"required,min=5,max=15"`
	Password string	`json:"password" form:"password" binding:"required,min=5,max=15"`
	Email string	`json:"email" form:"email" binding:"required,min=2,max=15" binding:"required,min=8,max=16"`
	Nickname string	`json:"nickname" form:"nickname" binding:"required,min=2,max=10"`
}

type AdminLoginService struct {
	Username string	`json:"username" form:"username" binding:"required,min=5,max=15"`
	Password string	`json:"password" form:"password" binding:"required,min=5,max=15"`
}

//管理员注册
func (service *AdminRegisterService) AdminRegister() *serializer.Response {
	var user model.Admin
	var count int
	model.DB.Model(&user).Where("username=?",service.Username).Count(&count)
	if	count == 1{
		return &serializer.Response{
			Status: 400,
			Msg: "账号已存在",
		}
	}

	if err:=model.DB.Model(&user).Where("nickname=?",service.Nickname).Error;err!=nil{
		return &serializer.Response{
			Status: 400,
			Msg: "该昵称已存在",
		}
	}
	if err:=model.DB.Model(&user).Where("email=?",service.Email).Error;err!=nil{
		return &serializer.Response{
			Status: 400,
			Msg: "该邮箱已注册用户",
		}
	}
	user.Username=service.Username
	user.Email=service.Email
	user.Nickname=service.Nickname

	if err:=user.SetAdminpassword(service.Password);err!=nil{
		logging.Info(err)
		return &serializer.Response{
			Status:400,
			Msg: "加密失败",
		}
	}
	if err:=model.DB.Create(&user).Error;err!=nil{
		logging.Info(err)
		return &serializer.Response{
			Status: 400,
			Msg: "注册失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg: "注册成功",

	}
}

//管理员登录
func (service *AdminLoginService) AdminLogin()serializer.Response  {
	var user model.Admin
	if err:=model.DB.Where("username=?",service.Username).First(&user).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			return serializer.Response{
				Status: 400,
				Msg: "账号不存在",
			}
		}
		return serializer.Response{
			Status: 400,
			Msg: "数据库错误",
		}
	}

	if user.CheckAdminpassword(service.Password) ==false{
		return serializer.Response{
			Status: 400,
			Msg: "密码错误",
		}
	}
	token,err:=utils.Generatetoken(user.ID,service.Username,service.Password,1)
	fmt.Println(service.Username,service.Password)
	if err != nil {
		logging.Info(err)
		return serializer.Response{
			Status: 400,
			Msg: "签发token失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{User: serializer.Builduadmin(&user),Token:token,},
		Msg: "登陆成功",
	}

}