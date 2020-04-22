package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 跨域处理中间件
func CrossDomain(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,Admin-Token")
	c.Header("Access-Control-Allow-Methods", "GET,PUT,PATCH,POST,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Content-Type", "application/json")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
