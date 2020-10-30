package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/z1px/framework/conf/env"
	"github.com/z1px/framework/util/filepath"
	"gopkg.in/ini.v1"
	"log"
	"path"
)

// 服务配置
type serverConf struct {
	Name string `toml:"name"`
	Mode string `toml:"mode"`
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// MYSQL数据库配置
type mysqlConf struct {
	Driver    string `toml:"driver"`
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	Database  string `toml:"database"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Charset   string `toml:"charset"`
	Collation string `toml:"collation"`
	Prefix    string `toml:"prefix"`
	Debug     bool   `toml:"debug"`
}

// REDIS数据库配置
type redisConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

// 基础配置
type baseConf struct {
	Server serverConf `toml:"server"`
	Mysql  mysqlConf  `toml:"mysql"`
	Redis  redisConf  `toml:"redis"`
}

var (
	Base *baseConf // 基础配置
	App  *ini.File // 通用配置
)

// 初始化配置文件
func Init() {
	// 加载TOML配置
	LoadTomlConf()
	// 加载ENV配置
	LoadEnvConf()
	// 加载INI配置
	LoadIniConf()
}

// 加载TOML基础配置文件
func LoadTomlConf() {
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// 公共配置文件完整路径
	filename := path.Join(confPath, "base.toml")
	// 判断配置文件是否存在
	if filepath.IsExist(filename) {
		if _, err := toml.DecodeFile(filename, &Base); err != nil {
			// handle error
			log.Fatalln(err)
		}
	}
}

// 加载ENV基础配置
func LoadEnvConf() {
	// 服务配置
	serverConf := &Base.Server
	serverConf.Name = env.GetString("APP_NAME", serverConf.Name)
	serverConf.Mode = env.GetString("APP_MODE", serverConf.Mode)
	serverConf.Host = env.GetString("APP_HOST", serverConf.Host)
	serverConf.Port = env.GetInt("APP_PORT", serverConf.Port)

	// mysql数据库配置
	mysqlConf := &Base.Mysql
	mysqlConf.Driver = env.GetString("DB_DRIVER", mysqlConf.Driver)
	mysqlConf.Host = env.GetString("DB_HOST", mysqlConf.Host)
	mysqlConf.Port = env.GetInt("DB_PORT", mysqlConf.Port)
	mysqlConf.Database = env.GetString("DB_DATABASE", mysqlConf.Database)
	mysqlConf.Username = env.GetString("DB_USERNAME", mysqlConf.Username)
	mysqlConf.Password = env.GetString("DB_PASSWORD", mysqlConf.Password)
	mysqlConf.Charset = env.GetString("DB_CHARSET", mysqlConf.Charset)
	mysqlConf.Collation = env.GetString("DB_COLLATION", mysqlConf.Collation)
	mysqlConf.Prefix = env.GetString("DB_PREFIX", mysqlConf.Prefix)
	mysqlConf.Debug = env.GetBool("DB_DEBUG", mysqlConf.Debug)

	// redis数据库配置
	redisConf := &Base.Redis
	redisConf.Host = env.GetString("REDIS_HOST", redisConf.Host)
	redisConf.Port = env.GetInt("REDIS_PORT", redisConf.Port)
	redisConf.Password = env.GetString("REDIS_PASSWORD", redisConf.Password)
	redisConf.Db = env.GetInt("REDIS_DB", redisConf.Db)
}

// 初始化INI配置
func LoadIniConf() {
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// 公共配置文件完整路径
	filename := path.Join(confPath, "app.ini")
	// 判断配置文件是否存在
	if filepath.IsExist(filename) {
		cfg, err := ini.Load(filename)
		if err != nil {
			// handle error
			log.Fatalln(err)
		} else {
			App = cfg
		}
	}
}

// 获取运行模式
func GetMode() (mode string) {
	mode = Base.Server.Mode
	return
}
