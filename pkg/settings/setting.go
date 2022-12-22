package settings

import (
	"github.com/spf13/viper"
	"log"
	"reflect"
	"time"
)

type MySQL struct {
	ConnMaxLifetime int
	MaxIdleConn     int
	MaxOpenConn     int
	ReadConfig      MySQLRead
	WriteConfig     MySQLWrite
}

type MySQLRead struct {
}

type MySQLWrite struct {
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
	loadBaseConfig(config)
}

// loadBaseConfig 加载基础配置
func loadBaseConfig(config *Config) {
	config.RunMode = viper.GetString("Application.RunMode")
	config.HTTPPort = viper.GetString("Server.Port")
	config.ReadTimeout = time.Duration(viper.GetInt64("Server.ReadTimeout")) * time.Second
	config.WriteTimeout = time.Duration(viper.GetInt64("Server.WriteTimeout")) * time.Second
}

// GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}

// GetValue 获取某个具体的配置值
func (c *Config) GetValue(key string) string {
	rv := reflect.ValueOf(c)
	rs := rv.Elem().FieldByName(key)

	if rs.IsValid() {
		return rs.String()
	}
	return ""
}
