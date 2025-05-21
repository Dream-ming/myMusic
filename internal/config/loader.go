package config

import (
	"github.com/spf13/viper"
)

func LoadOSSConfig() {
	viper.SetConfigName("oss")
	viper.AddConfigPath("./configs")
	viper.MergeInConfig()
}

func LoadMySQLConfig() {
	viper.SetConfigName("sql")
	viper.AddConfigPath("./configs")
	viper.MergeInConfig()
}

func LoadJWTConfig() {
	viper.SetConfigName("jwt")
	viper.AddConfigPath("./configs")
	viper.MergeInConfig()
}