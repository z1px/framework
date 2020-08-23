package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/z1px/framework/util/filepath"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

// 修改gin自带Logger
func Logger() gin.HandlerFunc {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	// gin.DisableConsoleColor()

	// 日志文件存在放目录
	logPath, _ := filepath.GetLogPath()
	// 日志文件完整路径
	logFile := path.Join(logPath, "gin.log")

	// 检测文件是否存在
	if filepath.IsExist(logFile) {
		// 备份旧日志文件
		err := os.Rename(logFile, path.Join(logPath, fmt.Sprintf(
			"gin.%d%02d%02d%02d%02d%02d.log",
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
	// 同时将日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(fp, os.Stdout)

	// 错误日志文件完整路径
	errorFile := path.Join(logPath, "error.log")
	// 错误日志记录到文件
	_fp, _err := os.OpenFile(errorFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if _err != nil {
		log.Fatalln("Failed to open error logger file：", _err)
	}
	gin.DefaultErrorWriter = io.MultiWriter(_fp, os.Stderr)

	// 定义路由日志的格式
	logFormatter := func(param gin.LogFormatterParams) (format string) {
		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		// 自定义日志格式
		format = fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\" %s \"%s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
		return
	}

	// 返回日志中间件
	return gin.LoggerWithFormatter(logFormatter)
}

// 日志写入文本
func LoggerToFile() gin.HandlerFunc {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	// gin.DisableConsoleColor()

	// 日志文件存在放目录
	logPath, _ := filepath.GetLogPath()
	// 日志文件完整路径
	logFile := path.Join(logPath, "logger.log")
	// 写入文件
	fp, _ := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	// 实例化
	Log := logrus.New()

	// 设置输出
	Log.SetOutput(fp)

	// 设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	// 设置日志格式为json格式
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, err := rotateLogs.New(
		// 分割后的文件名称
		strings.TrimSuffix(logFile, "log")+"%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(logFile),
		// 设置最大保存时间(7天)
		rotateLogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Fatalln("err：", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 设置日志格式为json格式
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		Log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
