package dao

import (
	"com.mensssy.LabMS/dao/db"
	"com.mensssy.LabMS/model"
)

func GetToken(userId string, tokenType string) (string, error) {
	db := db.SqlDB

	var col string
	if tokenType == "Mobile" {
		col = "token_mobile"
	} else if tokenType == "PC" {
		col = "token_pc"
	}

	var token string
	res := db.Model(&model.UserSecurity{}).Select(col).Where("user_id = ?", userId).First(&token)
	if res.Error != nil {
		return "", res.Error
	} else {
		return token, nil
	}
}

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
