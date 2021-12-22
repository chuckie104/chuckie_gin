package initialize

import (
	"chuckie_gin/global"
	"chuckie_gin/utils"
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {

	cfg := zap.NewDevelopmentConfig()

	cfg.OutputPaths = []string{
		fmt.Sprintf("%slog_%s.log", global.Settings.LogsAddress, utils.GetNowFormatTodayTime()), //
		"stdout",
	}

	logg, _ := cfg.Build()

	zap.ReplaceGlobals(logg)

	global.Lg = logg
}
