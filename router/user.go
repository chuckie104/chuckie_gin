package router

import (
	"chuckie_gin/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter (Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		UserRouter.POST("login", controller.PasswordLogin)
		UserRouter.GET("userlist", controller.GetUserList)
	}
}