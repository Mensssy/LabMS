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

func GetSecurityInfo(userId string) (*model.UserSecurity, error) {
	db := db.SqlDB

	var userSecurityInfo model.UserSecurity
	res := db.Model(&model.UserSecurity{}).Where("user_id = ?", userId).First(&userSecurityInfo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userSecurityInfo, nil
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

func CreateUser(info model.UserSecurity) error {
	db := db.SqlDB
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	//创建用户
	res := db.Create(&model.User{
		UserId:   info.UserId,
		UserName: "BaconSandwich",
	})
	if res.Error != nil {
		db.Rollback()
		return res.Error
	}
	//创建安全信息
	res = db.Create(&info)
	if res.Error != nil {
		db.Rollback()
		return res.Error
	}

	return tx.Commit().Error
}
