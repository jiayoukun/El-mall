package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Username string
	Name string
	Phone string
	Address string
}
