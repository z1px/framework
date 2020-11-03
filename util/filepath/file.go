package filepath

import (
	"fmt"
	"github.com/z1px/framework/conf/env"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 获取项目运行根目录
func GetRootPath() (rootPath string) {
	// 项目根目录
	rootPath, _ = os.Getwd()
	return
}

// 获取项目配置目录
func GetConfPath() (confPath string, err error) {
	// 配置文件目录
	confPath = path.Join(GetRootPath(), path.Join(env.GetString("CONF_PATH", "conf")))
	if !IsExist(confPath) {
		err = fmt.Errorf("文件夹不存在")
	}
	return
}

// 获取日志存放目录，不存在时创建
func GetLogPath() (logPath string, err error) {
	logPath, err = CreatePathIfNotExist(env.GetString("LOG_PATH", "runtime/log"))
	return
}

// 获取文件上传目录，不存在时创建
func GetUploadPath() (UploadPath string, err error) {
	UploadPath, err = CreatePathIfNotExist(env.GetString("UPLOAD_PATH", "static/upload"))
	return
}

// 获取文件目录，当目录不存在的时候递归创建
func CreatePathIfNotExist(fp string) (filePath string, err error) {
	// 完整文件目录
	filePath = path.Join(GetRootPath(), fp)
	// 检测文件夹是否存在
	if !IsExist(filePath) {
		// 必须分成两步
		// 先创建文件夹
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			return
		}
		// 再修改权限
		if err = os.Chmod(filePath, os.ModePerm); err != nil {
			return
		}
	}
	return
}

// 检测文件夹或者文件是否存在是否存在
func IsExist(filename string) bool {
	// 检测文件夹是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// 获取文件夹下所有文件
func GetPathFiles(path string) (list []string) {
	// 检测文件夹是否存在
	if IsExist(path) {
		// 读取目录下文件
		files, err := ioutil.ReadDir(path)
		if err == nil {
			for _, file := range files {
				if file.IsDir() { // 跳过目录
					continue
				}
				list = append(list, file.Name())
			}
		}
	}
	return
}

// 获取文件名，不带后缀
func getFileName(filename string) (name string) {
	name = strings.TrimSuffix(path.Base(filename), path.Ext(filename))
	return
}
