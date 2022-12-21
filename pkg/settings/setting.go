package settings

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type MySQL struct {
}

type Redis struct {
}

type Config struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MySQL        MySQL
	Redis        Redis
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

	config = &Config{}
	config.loadBaseConfig()
}

// loadBaseConfig 加载基础配置
func (c *Config) loadBaseConfig() {
	config.RunMode = viper.GetString("Application.RunMode")
	config.HTTPPort = viper.GetString("Server.Port")
	config.ReadTimeout = time.Duration(viper.GetInt64("Server.ReadTimeout")) * time.Second
	config.WriteTimeout = time.Duration(viper.GetInt64("Server.WriteTimeout")) * time.Second
}

// GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}
