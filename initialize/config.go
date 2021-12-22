package initialize

import (
	"chuckie_gin/config"
	"chuckie_gin/global"
	"github.com/spf13/viper"
)

func Initialize() {
	v := viper.New()

	v.SetConfigFile("./settings-dev.yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := config.ServerConfig{}

	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	global.Settings = serverConfig
}
