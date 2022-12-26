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
	ReadConfig      *MySQLConnConfig
	WriteConfig     *MySQLConnConfig
}

type MySQLConnConfig struct {
	Addr   string
	DBName string
	User   string
	Passwd string
}

type Redis struct {
}

type Config struct {
	RunMode      string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MySQL        *MySQL
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
	loadMySQLConfig(config)
}

// loadBaseConfig 加载基础配置
func loadBaseConfig(config *Config) {
	config.RunMode = viper.GetString("Application.RunMode")
	config.HTTPPort = viper.GetString("Server.Port")
	config.ReadTimeout = time.Duration(viper.GetInt64("Server.ReadTimeout")) * time.Second
	config.WriteTimeout = time.Duration(viper.GetInt64("Server.WriteTimeout")) * time.Second
}

// loadMySQLConfig 加载 MySQL 相关配置
func loadMySQLConfig(config *Config) {

	var MySQLConfig = new(MySQL)

	MySQLConfig.ConnMaxLifetime = viper.GetInt("MySQL.Base.ConnMaxLifetime")
	MySQLConfig.MaxIdleConn = viper.GetInt("MySQL.Base.MaxIdleConn")
	MySQLConfig.MaxOpenConn = viper.GetInt("MySQL.Base.MaxOpenConn")

	var ReadMySQLConfig = new(MySQLConnConfig)
	ReadMySQLConfig.Addr = viper.GetString("MySQL.Read.Addr")
	ReadMySQLConfig.DBName = viper.GetString("MySQL.Read.DBName")
	ReadMySQLConfig.User = viper.GetString("MySQL.Read.User")
	ReadMySQLConfig.Passwd = viper.GetString("MySQL.Read.Passwd")

	var WriteMySQLConfig = new(MySQLConnConfig)
	WriteMySQLConfig.Addr = viper.GetString("MySQL.Write.Addr")
	WriteMySQLConfig.DBName = viper.GetString("MySQL.Write.DBName")
	WriteMySQLConfig.User = viper.GetString("MySQL.Write.User")
	WriteMySQLConfig.Passwd = viper.GetString("MySQL.Write.Passwd")

	MySQLConfig.ReadConfig = ReadMySQLConfig
	MySQLConfig.WriteConfig = WriteMySQLConfig

	config.MySQL = MySQLConfig

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
