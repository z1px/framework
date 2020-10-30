package framework

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/z1px/framework/conf"
	_ "github.com/z1px/framework/db/ip"
	"github.com/z1px/framework/db/mysql"
	"github.com/z1px/framework/db/redis"
	"github.com/z1px/framework/logs"
	"github.com/z1px/framework/middleware"
	"github.com/z1px/framework/router"
	"net/http"
	"time"
)

// 初始化
func Init() (engine *gin.Engine) {
	// 初始化配置文件
	conf.Init()

	// 设置运行模式
	gin.SetMode(conf.GetMode())

	// 新建一个没有任何默认中间件的路由
	engine = gin.New()
	// 全局中间件
	// 注册日志记录器中间件
	engine.Use(logs.Logger())
	// Engine.Use(logger.LoggerToFile())
	// 注册崩溃恢复中间件，异常捕获中间件
	engine.Use(gin.Recovery(), middleware.Recover())
	// 崩溃恢复中间件
	engine.Use()

	// 初始化MYSQL数据库连接
	mysql.CreateDatabase()
	mysql.Connect()

	// 初始化REDIS数据库连接
	redis.Connect()

	// 初始化路由
	router.Init(engine)

	return
}

// 运行监听
func Run(engine *gin.Engine) {

	// 监听地址
	address := fmt.Sprintf("%s:%d", conf.Base.Server.Host, conf.Base.Server.Port)

	logs.Printf("listen：%s\n", address)

	// 自定义HTTP配置
	server := &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务并监听
	if err := server.ListenAndServe(); err != nil {
		logs.Fatal(err)
	}
}

// 关闭连接
func Close() {
	// 关闭mysql数据库连接
	mysql.Close()

	// 关闭redis连接
	redis.Close()
}
