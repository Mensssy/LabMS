package router

import (
	"time"

	"com.mensssy.LabMS/controller/response"
	"com.mensssy.LabMS/dao"
	"com.mensssy.LabMS/model"
	"com.mensssy.LabMS/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(getCors())

	r.POST("/test", func(c *gin.Context) {
		if err := dao.UpdateSecurityInfo(model.UserSecurity{
			UserId:      "1",
			TokenMobile: "1234",
		}); err != nil {
			c.JSON(response.Internal_Server_Error, response.Body{
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}

		c.JSON(response.OK, response.Body{
			Msg:  "test succeeded",
			Data: nil,
		})
	})

	api := r.Group("/api")
	{
		api.POST("/login", service.Login)
		user := api.Group("/users")
		{
			user.GET("")
		}

	}

	return r
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
