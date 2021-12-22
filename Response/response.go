package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success (context *gin.Context, code int, msg interface{}, data interface{},) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg": msg,
		"data": data,
	})
}


func Err (context *gin.Context, httpCode int ,code int, msg interface{}, data interface{},) {
	context.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg": msg,
		"data": data,
	})
}