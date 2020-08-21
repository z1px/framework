package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/z1px/framework/logs"
)

// 后置中间件
func After() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before request
		logs.DebugPrintln("After middleware before request")

		c.Next()

		// after request
		logs.DebugPrintln("After middleware after request")
	}
}
