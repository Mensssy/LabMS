package service

import (
	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"com.mensssy.LabMS/model"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	val, _ := c.Get("userId")
	userId, _ := val.(string)
	user, err := dao.GetUserInfo(userId)
	if err != nil {
		c.JSON(response.Internal_Server_Error, response.Body{
			Data: nil,
			Msg:  "userId not exists",
		})
	} else {
		c.JSON(response.OK, response.Body{
			Data: user,
			Msg:  "succeeded",
		})
	}
}

func UpdateUserInfo(c *gin.Context) {
	var userInfo model.User
	err := c.ShouldBindJSON(&userInfo)

	userId, _ := c.Get("userId")
	userInfo.UserId = userId.(string)
	userInfo.UserType = ""

	if err != nil {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "illegal user info structure",
		})
		return
	}

	err = dao.UpdateUserInfo(userInfo)
	if err != nil {
		c.AbortWithStatusJSON(response.Bad_Request, response.Body{
			Msg: "database error",
		})
		return
	}

	c.JSON(response.OK, response.Body{
		Msg: "update succeeded",
	})
}
