package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/z1px/framework/config"
	"os"
	"time"
)

// 判断是否是调试模式
func IsDebug() bool {
	return config.TomlConf.GetMode() == gin.DebugMode
}

// 格式化日志前缀
func FormatPrefix() (prefix string) {
	prefix = fmt.Sprintf(
		"[%d-%02d-%02d %02d:%02d:%02d]  ",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second())
	return
}

// 打印日志前缀
func PrintPrefix() {
	_, err := fmt.Fprint(gin.DefaultWriter, FormatPrefix())
	if err != nil {
		fmt.Println(err)
	}
}

// 打印错误日志前缀
func PrintErrPrefix() {
	_, err := fmt.Fprint(gin.DefaultErrorWriter, FormatPrefix())
	if err != nil {
		fmt.Println(err)
	}
}

func Println(a ...interface{}) {
	PrintPrefix()
	if _, err := fmt.Fprintln(gin.DefaultWriter, a...); err != nil {
		fmt.Println(err)
	}
}

func Print(a ...interface{}) {
	PrintPrefix()
	if _, err := fmt.Fprint(gin.DefaultWriter, a...); err != nil {
		fmt.Println(err)
	}
}

func Printf(format string, a ...interface{}) {
	PrintPrefix()
	if _, err := fmt.Fprintf(gin.DefaultWriter, format, a...); err != nil {
		fmt.Println(err)
	}
}

func DebugPrintln(a ...interface{}) {
	if IsDebug() {
		Println(a...)
	}
}

func DebugPrint(a ...interface{}) {
	if IsDebug() {
		Print(a...)
	}
}

func DebugPrintf(format string, a ...interface{}) {
	if IsDebug() {
		Printf(format, a...)
	}
}

func ErrPrintln(a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprintln(gin.DefaultErrorWriter, a...); err != nil {
		fmt.Println(err)
	}
}

func ErrPrint(a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprint(gin.DefaultErrorWriter, a...); err != nil {
		fmt.Println(err)
	}
}

func ErrPrintf(format string, a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprintf(gin.DefaultErrorWriter, format, a...); err != nil {
		fmt.Println(err)
	}
}

func Fatalln(a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprintln(gin.DefaultErrorWriter, a...); err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

func Fatal(a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprint(gin.DefaultErrorWriter, a...); err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	PrintErrPrefix()
	if _, err := fmt.Fprintf(gin.DefaultErrorWriter, format, a...); err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}
