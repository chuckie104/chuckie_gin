package controller

import (
	"chuckie_gin/Response"
	"chuckie_gin/dao"
	"chuckie_gin/forms"
	"chuckie_gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PasswordLogin (context *gin.Context) {
	passwordLoginForm := forms.PasswordLoginForm{}

	if err := context.ShouldBind(&passwordLoginForm); err != nil {
		utils.HandleValidatorError(context, err)
		return
	}
	Response.Success(context, 0, "success", nil)
}

func GetUserList (context *gin.Context) {
	userListForm := forms.UserListForm{}

	if err := context.ShouldBind(&userListForm); err != nil {
		utils.HandleValidatorError(context, err)
		return
	}

	total, userList := dao.GetUserListDao(userListForm.Page, userListForm.PageSize)

	if (total + len(userList)) == 0 {
		Response.Err(context, http.StatusBadRequest, http.StatusBadRequest, "未获取到数据", map[string]interface{}{
			"total": total,
			"data": userList,
			"page": userListForm.Page,
			"pageSize": userListForm.PageSize,
		})
		return
	}

	Response.Success(context,0, "获取用户列表成功", map[string]interface{}{
		"total": total,
		"data": userList,
		"page": userListForm.Page,
		"pageSize": userListForm.PageSize,
	})
}
