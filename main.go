package main

import (
	"chuckie_gin/global"
	"chuckie_gin/initialize"
	"chuckie_gin/middlewares"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
)

func main () {
	// 初始化配置
	initialize.Initialize()

	initialize.InitLogger()

	initialize.InitMysqlDB()

	r := initialize.Routers()

	r.Use(middlewares.GinLogger())

	color.Cyan("go-gin服务开始了")


	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	err := r.Run(fmt.Sprintf(":%d", global.Settings.Port))

	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}
}