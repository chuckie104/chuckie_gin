package utils

import (
	"chuckie_gin/Response"
	"chuckie_gin/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strings"
)

func HandleValidatorError (c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": err.Error(),
		//})
		Response.Err(c, http.StatusInternalServerError, 500, "字段校验错误", errs.Error())
	}
	msg := removeTopStruct(errs.Translate(global.Trans))
	Response.Err(c, http.StatusBadRequest, 400, "字段校验错误", msg)
	return
}

func removeTopStruct(fields map[string]string) map[string]string {
	rep := map[string]string{}

	for field, err := range fields {
		rep[field[strings.Index(field, ".")+1:]] = err
	}
	return rep
}

func ValidateMobile (fl validator.FieldLevel) bool {
	mobile := fl.Field().String()

	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)

	if !ok {
		return false
	}
	return true
}