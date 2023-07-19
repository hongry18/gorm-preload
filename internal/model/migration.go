package model

import "gorm.io/gorm"

func Migration(db *gorm.DB) {
	db.Exec("drop table orders")
	db.Exec("drop table users")
	if err := db.AutoMigrate(
		&User{},
		&Order{},
	); err != nil {
		panic(err)
	}
}
