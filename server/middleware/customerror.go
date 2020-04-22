package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 统一异常处理
func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if c.IsAborted() {
				c.Status(200)
				return
			}
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}

					c.Status(statusCode)
					fmt.Println(
						time.Now().Format("\n 2006-01-02 15:04:05.9999"),
						"[ERROR]",
						c.Request.Method,
						c.Request.URL,
						statusCode,
						c.Request.RequestURI,
						c.ClientIP(),
						p[2],
					)
					c.JSON(http.StatusOK, gin.H{
						"succ":true,
						"code": statusCode,
						"msg":  p[2],
					})
				}

			default:
				panic(err)
			}
		}
	}()
	c.Next()
}
