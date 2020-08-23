package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/z1px/framework/config"
	"github.com/z1px/framework/logs"
	"time"
)

// DB 数据库连接单例
var DB *gorm.DB

// 在中间件中初始化mysql连接
func Init() {
	// 获取数据库配置
	conf := config.DBConf.Mysql

	// 连接数据库
	db, err := gorm.Open(conf.Driver, fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset))
	// Error
	if err != nil {
		// handle error
		logs.ErrPrintf("mysql connect error：%v\n", err)
	}
	if db.Error != nil {
		logs.ErrPrintf("mysql error：%v\n", db.Error)
	}
	// 设置输出数据库日志
	db.LogMode(conf.Debug)
	db.SetLogger(gorm.Logger{LogWriter: NewLog()})
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	// 更改默认表名，设置数据库表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) (tableName string) {
		tableName = conf.Prefix + defaultTableName
		return
	}
	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(50)
	// 打开
	db.DB().SetMaxOpenConns(100)
	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}

// 检测数据库是否有连接
func IsOpen() bool {
	if err := DB.DB().Ping(); err != nil {
		logs.ErrPrintf("mysql ping error：%v\n", err)
		return false
	} else {
		return true
	}
}

// 关闭连接
func Close() {
	err := DB.Close()
	if err != nil {
		logs.ErrPrintf("mysql close error：%v\n", err)
	}
}
