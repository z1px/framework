package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/z1px/framework/config"
	"github.com/z1px/framework/config/env"
	"github.com/z1px/framework/util/filepath"
	"log"
	"path"
)

// 初始化TOML配置
func Init() {
	// 加载公共配置文件
	LoadConfToml()
	// 加载数据库配置文件
	LoadDatabaseToml()

	// 加载公共ENV配置
	env.LoadConfEnv()
	// 加载数据库ENV配置
	env.LoadDatabaseEnv()
}

// 加载公共配置文件
func LoadConfToml() {
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// 公共配置文件完整路径
	filename := path.Join(confPath, "config.toml")
	// 判断配置文件是否存在
	if filepath.IsExist(filename) {
		if _, err := toml.DecodeFile(filename, &config.TomlConf); err != nil {
			// handle error
			log.Fatalln(err)
		}
	}
}

// 加载数据库配置文件
func LoadDatabaseToml() {
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// 数据库配置文件完整路径
	filename := path.Join(confPath, "database.toml")
	// 判断配置文件是否存在
	if filepath.IsExist(filename) {
		if _, err := toml.DecodeFile(filename, &config.DBConf); err != nil {
			// handle error
			log.Fatalln(err)
		}
	}
}
