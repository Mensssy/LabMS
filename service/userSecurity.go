package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"com.mensssy.LabMS/conf"
	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"com.mensssy.LabMS/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	userId := c.PostForm("userId")
	password := c.PostForm("password")
	device := c.PostForm("device")

	if device != "PC" && device != "Mobile" {
		//设备不支持
		c.JSON(response.Bad_Request, response.Body{
			Msg:  "device not supported",
			Data: nil,
		})
		return
	}

	eUserId := encrypt(userId)
	if securityInfo, err := dao.GetSecurityInfo(eUserId); err != nil {
		//用户不存在
		c.JSON(response.Internal_Server_Error, response.Body{
			Msg:  "userId not exists",
			Data: nil,
		})
	} else {
		//密码加盐哈希
		ePassword := encrypt(password + securityInfo.Salt)

		if ePassword == securityInfo.Password {
			//获取token
			token := generateToken(userId, device)

			//更新token
			if device == "PC" {
				dao.UpdateSecurityInfo(model.UserSecurity{
					UserId:  eUserId,
					TokenPC: token,
				})
			} else if device == "Mobile" {
				dao.UpdateSecurityInfo(model.UserSecurity{
					UserId:      eUserId,
					TokenMobile: token,
				})
			}

			//返回token
			c.JSON(response.OK, response.Body{
				Msg: "login succeeded",
				Data: map[string]interface{}{
					"token":       token,
					"tokenExTime": time.Now().Add(time.Hour * 2).Unix(),
				}})

		} else {
			//密码错误
			c.JSON(response.Unauthorized, response.Body{
				Msg:  "wrong password",
				Data: nil,
			})
		}
	}
}

func encrypt(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}

func generateToken(userId string, device string) string {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
		"id":   userId,
		"type": device,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenStr, err := token.SignedString([]byte(conf.TokenKey)); err != nil {
		return "error"
	} else {
		return tokenStr
	}
}
