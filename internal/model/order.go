package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

type Orders []Order
