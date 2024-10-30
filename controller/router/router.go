package router

import (
	"time"

	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/service"
	"com.mensssy.LabMS/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(getCors())

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(response.OK, response.Body{
			Msg:  "U GOT ME!!",
			Data: nil,
		})
	})

	api := r.Group("/api")
	{
		api.POST("/login", service.Login)
		api.POST("/signin", service.Signin)

		common := api.Group("")
		//token鉴权后上下文中将存储userId
		common.Use(tokenAuth())
		user := common.Group("/users")
		{
			user.GET("", service.GetUserInfo)
		}

	}

	return r
}

func tokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		msg, err := util.ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(response.Bad_Request, response.Body{
				Data: nil,
				Msg:  err.Error(),
			})
			return
		} else if msg == "invalid_token" {
			c.AbortWithStatusJSON(response.Unauthorized, response.Body{
				Data: nil,
				Msg:  "invalid token",
			})
			return
		}

		//成功 保存userId，方便后续操作
		c.Set("userId", msg)
		c.Next()
	}
}

func getCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
