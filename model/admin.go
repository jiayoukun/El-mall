package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"unique"`
	PasswordDigest string
	Email string `gorm:"unique"`
	Nickname string `gorm:"unique"`
}

//加密密码
func (user *Admin)SetAdminpassword(password string) error {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {

		return err
	}
	user.PasswordDigest=string(bytes)
	return nil
}
//校验密码
func (user *Admin)CheckAdminpassword(password string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest),[]byte(password))

	return err==nil
}
