package db

import "com.mensssy.LabMS/model"

func SqlMigrate() {
	db := SqlDB

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserSecurity{})
	db.AutoMigrate(&model.Invoice{})

}
