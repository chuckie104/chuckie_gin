package config

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
	LogsAddress string `mapstructure:"logsAddress"`
	Mysqlinfo   MysqlConfig `mapstructure:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}