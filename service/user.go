package service

import (
	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	val, _ := c.Get("userId")
	userId, _ := val.(string)
	user, err := dao.GetUserInfo(userId)
	if err != nil {
		c.JSON(response.Internal_Server_Error, response.Body{
			Data: nil,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(response.OK, response.Body{
			Data: user,
			Msg:  "succeeded",
		})
	}
}
