package forms

type PasswordLoginForm struct {
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Username string `form:"name" json:"name" binding:"required,mobile"`
}

type UserListForm struct {
	Page int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}