package env

import (
	"github.com/joho/godotenv"
	"github.com/z1px/framework/config"
	"log"
	"os"
	"strconv"
)

// 初始化ENV配置文件
func Init() {
	// 读取配置
	if err := godotenv.Load(); err != nil {
		// handle error
		log.Fatal(err)
	}
}

// 获取ENV配置string类型
func GetString(key string, value string) string {
	val, exist := os.LookupEnv(key)
	if exist {
		value = val
	}
	return value
}

// 获取ENV配置int类型
func GetInt(key string, value int) int {
	val, exist := os.LookupEnv(key)
	if exist {
		value, _ = strconv.Atoi(val)
	}
	return value
}

// 获取ENV配置bool类型
func GetBool(key string, value bool) bool {
	val, exist := os.LookupEnv(key)
	if exist {
		value, _ = strconv.ParseBool(val)
	}
	return value
}

// 加载公共配置文件
func LoadConfEnv() {
	// 应用配置
	appConf := &config.TomlConf.App
	appConf.Name = GetString("APP_NAME", appConf.Name)
	appConf.Mode = GetString("APP_MODE", appConf.Mode)

	// 服务配置
	serverConf := &config.TomlConf.Server
	serverConf.Host = GetString("APP_HOST", serverConf.Host)
	serverConf.Port = GetInt("APP_PORT", serverConf.Port)
}

// 加载数据库配置文件
func LoadDatabaseEnv() {
	// mysql数据库配置
	mysqlConf := &config.DBConf.Mysql
	mysqlConf.Driver = GetString("DB_DRIVER", mysqlConf.Driver)
	mysqlConf.Host = GetString("DB_HOST", mysqlConf.Host)
	mysqlConf.Port = GetInt("DB_PORT", mysqlConf.Port)
	mysqlConf.Database = GetString("DB_DATABASE", mysqlConf.Database)
	mysqlConf.Username = GetString("DB_USERNAME", mysqlConf.Username)
	mysqlConf.Password = GetString("DB_PASSWORD", mysqlConf.Password)
	mysqlConf.Charset = GetString("DB_CHARSET", mysqlConf.Charset)
	mysqlConf.Collation = GetString("DB_COLLATION", mysqlConf.Collation)
	mysqlConf.Prefix = GetString("DB_PREFIX", mysqlConf.Prefix)
	mysqlConf.Debug = GetBool("DB_DEBUG", mysqlConf.Debug)

	// redis数据库配置
	redisConf := &config.DBConf.Redis
	redisConf.Host = GetString("REDIS_HOST", redisConf.Host)
	redisConf.Port = GetInt("REDIS_PORT", redisConf.Port)
	redisConf.Password = GetString("REDIS_PASSWORD", redisConf.Password)
	redisConf.Db = GetInt("REDIS_DB", redisConf.Db)
}
