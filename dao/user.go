package dao

import (
	"com.mensssy.LabMS/dao/db"
	"com.mensssy.LabMS/model"
)

func GetUserInfo(userId string) (*model.User, error) {
	db := db.SqlDB

	var user model.User
	res := db.Model(&model.User{}).Omit("user_id").Where("user_id = ?", userId).First(&user)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return &user, nil
	}
}
