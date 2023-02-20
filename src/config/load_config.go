package config

import "github.com/spf13/viper"

var AppConfig = viper.New()

func init() {
	AppConfig.AddConfigPath("../../")
	AppConfig.SetConfigName("application")
	AppConfig.SetConfigType("yml")

	if err := AppConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}
