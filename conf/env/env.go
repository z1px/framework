package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// 初始化ENV配置文件
func init() {
	// 读取配置
	if err := godotenv.Load(); err != nil {
		// handle error
		log.Fatalln(err)
	}
}

// 获取ENV配置string类型
func GetString(key string, value string) string {
	val, exist := os.LookupEnv(key)
	if exist {
		value = val
	}
	return value
}

// 获取ENV配置int类型
func GetInt(key string, value int) int {
	val, exist := os.LookupEnv(key)
	if exist {
		value, _ = strconv.Atoi(val)
	}
	return value
}

// 获取ENV配置bool类型
func GetBool(key string, value bool) bool {
	val, exist := os.LookupEnv(key)
	if exist {
		value, _ = strconv.ParseBool(val)
	}
	return value
}
