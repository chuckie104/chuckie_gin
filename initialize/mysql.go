package initialize

import (
	"chuckie_gin/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB () {
	mysqlInfo := global.Settings.Mysqlinfo

	// 拼接地址
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DbName)
	db, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 连接数据库出错记录日志
	if err != nil {
		zap.L().Info(err.Error())
	}

	global.DB = db
}
