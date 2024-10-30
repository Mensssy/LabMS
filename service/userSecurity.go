package service

import (
	"time"

	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"com.mensssy.LabMS/model"
	"com.mensssy.LabMS/util"
	"github.com/gin-gonic/gin"
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

	eUserId := util.Encrypt(userId)
	if securityInfo, err := dao.GetSecurityInfo(eUserId); err != nil {
		//用户不存在
		c.JSON(response.Internal_Server_Error, response.Body{
			Msg:  "user not exists",
			Data: nil,
		})
	} else {
		//密码加盐哈希
		ePassword := util.Encrypt(password + securityInfo.Salt)

		if ePassword == securityInfo.Password {
			//获取token
			token := util.GenerateToken(userId, device)

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

func Signin(c *gin.Context) {
	userId := c.PostForm("userId")
	password := c.PostForm("password")

	salt := util.GetSalt()

	eUserId := util.Encrypt(userId)
	ePassword := util.Encrypt(password + salt)

	err := dao.CreateUser(model.UserSecurity{
		UserId:   eUserId,
		Password: ePassword,
		Salt:     salt,
	})

	if err != nil {
		c.JSON(response.Internal_Server_Error, response.Body{
			Msg:  "signin failed",
			Data: err.Error(),
		})
	} else {
		c.JSON(response.OK, response.Body{
			Msg:  "signin succeeded",
			Data: nil,
		})
	}
}
