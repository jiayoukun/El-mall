package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	PasswordDigest string //密文
	Nickname string `gorm:"unique"`//昵称
}
//加密密码
func (user *User)Setpassword(password string) error {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),15)
	if err != nil {

		return err
	}
	user.PasswordDigest=string(bytes)
	return nil
}
//校验密码
func (user *User)Checkpassword(password string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))

	return err==nil
}