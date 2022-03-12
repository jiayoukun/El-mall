package model

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	Username string
	ProductID string
	Num string
	MaxNum string
	ProductName string
	Imgpath string
	Price string

}
