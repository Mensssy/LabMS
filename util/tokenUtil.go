package util

import (
	"time"

	"com.mensssy.LabMS/conf"
	"com.mensssy.LabMS/dao"
	"github.com/golang-jwt/jwt/v5"
)

type myClaims struct {
	jwt.RegisteredClaims
	Id        string `json:"Id"`
	TokenType string `json:"Tokentype"`
}

func GenerateToken(userId string, device string) string {
	claims := myClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
		Id:        userId,
		TokenType: device,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenStr, err := token.SignedString([]byte(conf.TokenKey)); err != nil {
		return "error"
	} else {
		return tokenStr
	}
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.TokenKey), nil
	})

	//token解析错误
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*myClaims); token.Valid && ok {
		eid := Encrypt(claims.Id)

		trueTokenStr, err := dao.GetToken(eid, claims.TokenType)

		//数据库出错
		if err != nil {
			return "", err
		} else if trueTokenStr != tokenStr { //该设备token失效
			return "invalid_token", nil
		}

		return eid, nil
	}

	return "invalid_token", nil
}
