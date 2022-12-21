package settings

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func newConfig() *Config {
	return &Config{}
}

var config *Config

func init() {

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig() //读取
	if err != nil {
		log.Fatalln("read config failed: %v", err)
	}

	config = newConfig()
	config.RunMode = viper.GetString("Application.RunMode")
	config.HTTPPort = viper.GetString("Server.Port")
	config.ReadTimeout = time.Duration(viper.GetInt64("Server.ReadTimeout")) * time.Second
	config.WriteTimeout = time.Duration(viper.GetInt64("Server.WriteTimeout")) * time.Second

}

func GetConfig() *Config {
	return config
}
