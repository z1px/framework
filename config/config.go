package config

import (
	"gopkg.in/ini.v1"
)

// 应用配置
type AppConf struct {
	Name string `toml:"name"`
	Mode string `toml:"mode"`
}

// 服务配置
type ServerConf struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// 公共配置
type ConfigConf struct {
	App    AppConf    `toml:"app"`
	Server ServerConf `toml:"server"`
}

// MYSQL数据库配置
type MysqlConf struct {
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
type RedisConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

// 数据库配置
type DatabaseConf struct {
	Mysql MysqlConf `toml:"mysql"`
	Redis RedisConf `toml:"redis"`
}

// 初始化数据库配置结构体
var TomlConf *ConfigConf

// 初始化数据库配置结构体
var DBConf *DatabaseConf

// 初始化INI配置结构体
var IniConf *ini.File

// 获取运行模式
func (conf *ConfigConf) GetMode() (mode string) {
	mode = conf.App.Mode
	return
}
