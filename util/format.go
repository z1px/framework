package util

import (
	"regexp"
	"strings"
)

// 获取客户端
func Client(userAgent string) (client string) {
	userAgent = strings.ToLower(userAgent)
	reg, _ := regexp.Compile("(?i:Mobile|iPod|iPhone|Android|Opera Mini|BlackBerry|webOS|UCWEB|Blazer|PSP)")
	client = reg.FindString(userAgent)
	return
}
