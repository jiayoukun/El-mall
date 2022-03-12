package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Username string `gorm:"primary_key`
	ProductID string
	SumPrice string
	Num string
	OrderNum  string `gorm:"primary_key`
	Address string
	AddressName string
	AddressPhone string
}
