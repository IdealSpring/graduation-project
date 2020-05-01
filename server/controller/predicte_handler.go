package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostFileDownload(c *gin.Context) {
	content := "中国中国"

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=data.zip")
	c.Header("Content-type", "application/octet-stream")
	c.Writer.Write([]byte(content))
}
