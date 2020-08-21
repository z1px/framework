package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/z1px/framework/logs"
)

// 前置中间件
func Before() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		logs.DebugPrintln("Before middleware before request")

		c.Next()

		// after request
		logs.DebugPrintln("Before middleware after request")
	}
}
