package initialize

import (
	"chuckie_gin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.New()

	ApiGroup := r.Group("v1")

	router.UserRouter(ApiGroup)

	return r
}