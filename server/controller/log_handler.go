package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/server/common/utils"
	bind "graduation-project/server/controller/bind_param"
	"graduation-project/server/models"
	"net/http"
	"strconv"
)

var l = new(models.OperationLog)

// 删除所有操作日志
func DeleteDeleteAllOperationLog(c *gin.Context) {
	err := l.DeleteAllOperationLog()
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 删除所有登录日志
func DeleteDeleteAllLoginLog(c *gin.Context) {
	err := l.DeleteAllLoginLog()
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 删除单条日志
func DeleteOperationLog(c *gin.Context) {
	id := c.Param("logId")
	logId, _ := strconv.Atoi(id)

	err := l.DeleteOperationLog(logId)
	utils.AssertErr(err)

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

// 获取日志 to 列表
func PostFetchDataToPage(c *gin.Context) {
	var logParam bind.LogParam
	c.Bind(&logParam)

	isLoginOrLogout := logParam.IsLoginOrLogout
	moduleName := ""
	username := ""
	nick := ""
	current := 1
	size := 10

	if logParam.ModuleName != "" {
		moduleName = logParam.ModuleName
	}
	if logParam.Username != "" {
		username = logParam.Username
	}
	if logParam.Nick != "" {
		nick = logParam.Nick
	}
	if logParam.Current != 0 {
		current = logParam.Current
	}
	if logParam.Size != 0 {
		size = logParam.Size
	}

	roleList, total := l.FindLogPageByCondition(isLoginOrLogout, moduleName, username, nick, current, size)
	pages := total / size
	if total%size != 0 {
		pages++
	}

	var page = utils.Page{
		Records: roleList,
		Current: current,
		Pages:   pages,
		Size:    size,
		Total:   total,
	}

	c.JSON(http.StatusOK, gin.H{
		"succ": true,
		"page": page,
	})
}
