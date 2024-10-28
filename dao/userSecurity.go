package dao

import (
	"com.mensssy.LabMS/dao/db"
	"com.mensssy.LabMS/model"
)

func GetSecurityInfo(userId string) (model.UserSecurity, error) {
	db := db.SqlDB

	var userSecurityInfo model.UserSecurity
	res := db.Model(&model.UserSecurity{}).Where("user_id = ?", userId).First(&userSecurityInfo)
	if res.Error != nil {
		return userSecurityInfo, res.Error
	}
	return userSecurityInfo, nil
}

func UpdateSecurityInfo(info model.UserSecurity) error {
	db := db.SqlDB

	// res := db.Model(&info).Updates(&info) 等价 但下面的可读性更好
	res := db.Model(&model.UserSecurity{}).Where("user_id = ?", info.UserId).Updates(&info)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
