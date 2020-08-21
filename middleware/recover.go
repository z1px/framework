package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/z1px/framework/logs"
	"net/http"
	"time"
)

// 异常捕获中间件
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		logs.DebugPrintln("Recover middleware before request")

		defer func() {
			if r := recover(); r != nil {
				logs.ErrPrintln("recover：", r)
				c.JSON(http.StatusBadGateway, gin.H{
					"Code": 0,
					"Msg": r.(string),
					"Data": make(map[string]interface{}),
					"timestamp": time.Now().Unix(),
				})
			}
		}()

		c.Next()

		// after request
		logs.DebugPrintln("Recover middleware after request")
	}
}
