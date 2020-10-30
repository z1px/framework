package ip

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/z1px/framework/logs"
	"github.com/z1px/framework/util/filepath"
	"path"
)

// DB 数据库连接单例
var (
	City        *ipdb.City
	BaseStation *ipdb.BaseStation
	IDC         *ipdb.IDC
	District    *ipdb.District
	Language    = "CN"
)

// 初始化IP连接
func init() {
	InitBaseStation()
}

// 获取IP库文件地址
func getFilename() (filename string, err error) {
	// 获取配置文件夹完整路径
	confPath, _ := filepath.GetConfPath()
	// IP库文件完整路径
	filename = path.Join(confPath, "city.ipv4.ipdb")
	// 判断IP库文件是否存在
	if !filepath.IsExist(filename) {
		err = fmt.Errorf("IP库文件不存在，文件地址：%s\n", filename)
	}
	return
}

// 初始化连接
func InitCity() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err != nil {
		logs.ErrPrintln(err)
	}
	db, err := ipdb.NewCity(filename)
	if err != nil {
		// handle error
		logs.ErrPrintln(err)
	}

	City = db
}

func InitBaseStation() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err != nil {
		logs.ErrPrintln(err)
	}
	db, err := ipdb.NewBaseStation(filename)
	if err != nil {
		// handle error
		logs.ErrPrintln(err)
	}

	BaseStation = db
}

func InitIDC() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err != nil {
		logs.ErrPrintln(err)
	}
	db, err := ipdb.NewIDC(filename)
	if err != nil {
		// handle error
		logs.ErrPrintln(err)
	}

	IDC = db
}

func InitDistrict() { // 获取配置文件夹完整路径
	// 获取IP库文件地址
	filename, err := getFilename()
	if err != nil {
		logs.ErrPrintln(err)
	}
	db, err := ipdb.NewDistrict(filename)
	if err != nil {
		// handle error
		logs.ErrPrintln(err)
	}

	District = db
}

// 重新加载内容
func ReloadCity() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err == nil {
		// 更新 ipdb 文件后可调用 Reload 方法重新加载内容
		if err := City.Reload(filename); err != nil {
			logs.ErrPrintln(err)
		}
	}
}

func ReloadBaseStation() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err == nil {
		// 更新 ipdb 文件后可调用 Reload 方法重新加载内容
		if err := BaseStation.Reload(filename); err != nil {
			logs.ErrPrintln(err)
		}
	}
}

func ReloadIDC() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err == nil {
		// 更新 ipdb 文件后可调用 Reload 方法重新加载内容
		if err := IDC.Reload(filename); err != nil {
			logs.ErrPrintln(err)
		}
	}
}

func ReloadDistrict() {
	// 获取IP库文件地址
	filename, err := getFilename()
	if err == nil {
		// 更新 ipdb 文件后可调用 Reload 方法重新加载内容
		if err := District.Reload(filename); err != nil {
			logs.ErrPrintln(err)
		}
	}
}
