package ini

import (
	"github.com/z1px/framework/config"
	"github.com/z1px/framework/util/filepath"
	"gopkg.in/ini.v1"
	"log"
	"path"
)

// 初始化INI配置
func Init() {
	// 加载公共配置文件
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// 公共配置文件完整路径
	filename := path.Join(confPath, "app.ini")
	// 判断配置文件是否存在
	if filepath.IsExist(filename) {
		if cfg, err := ini.Load(filename); err != nil {
			// handle error
			log.Fatal(err)
		} else {
			config.IniConf = cfg
		}
	}
}
