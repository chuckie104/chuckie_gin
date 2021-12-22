package dao

import (
	"chuckie_gin/global"
	"chuckie_gin/models"
)

var users []models.User

func GetUserListDao (page int, page_size int) (int, []interface{}) {
	// 分页用户列表数据 长度为0  容量为len(users)
	userList := make([]interface{}, 0, len(users))

	// 计算偏移量
	offset := (page - 1) * page_size
	// 查询所有的user
	result := global.DB.Offset(offset).Limit(page_size).Find(&users)

	if result.RowsAffected == 0 {
		return 0, userList
	}

	// 获取users总数
	total := len(users)

	for _, userSingle := range users {
		birthday := ""
		if userSingle.Birthday == nil {
			birthday = ""
		} else {
			birthday = userSingle.Birthday.Format("2006-01-02")
		}

		userItemMap := map[string]interface{}{
			"id": userSingle.ID,
			"password": userSingle.Password,
			"name": userSingle.Name,
			"birthday": birthday,
		}
		userList = append(userList, userItemMap)
	}

	return total, userList
}