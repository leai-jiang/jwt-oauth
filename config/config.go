package config

import (
	"github.com/spf13/viper"
	"log"
)

var Viper *viper.Viper

func InitConfig() {
	Viper = viper.New()

	Viper.SetConfigName("config")
	Viper.AddConfigPath(".")
	Viper.SetConfigType("json")

	if err := Viper.ReadInConfig(); err != nil {
		log.Println("err: ", err)
	}
}
