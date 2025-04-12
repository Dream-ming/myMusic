package config

import (
	"github.com/spf13/viper"
)

func LoadOSSConfig() {
	viper.SetConfigName("oss")
	viper.AddConfigPath("./configs")
	viper.MergeInConfig()
}