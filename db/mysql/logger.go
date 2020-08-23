package mysql

import (
	"fmt"
	"github.com/z1px/framework/util/filepath"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func NewLog() *log.Logger {
	// 日志文件存在放目录
	logPath, _ := filepath.GetLogPath()
	// 日志文件完整路径
	logFile := path.Join(logPath, "sql.log")

	// 检测文件是否存在
	if filepath.IsExist(logFile) {
		// 备份旧日志文件
		err := os.Rename(logFile, path.Join(logPath, fmt.Sprintf(
			"sql.%d%02d%02d%02d%02d%02d.log",
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second())))
		if err != nil {
			log.Fatalln("Failed to rename error logger file：", err)
		}
	}

	// 记录到文件
	fp, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error logger file：", err)
	}

	//自定义日志格式
	return log.New(io.MultiWriter(fp, os.Stdout), "\r\n", 0)
}
