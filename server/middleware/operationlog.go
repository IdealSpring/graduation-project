package middleware

import (
	"github.com/gin-gonic/gin"
	"graduation-project/server/models"
	"strings"
	"time"
)

const (
	LOGIN  = "登录"
	LOGOUT = "登出"
	ADD    = "增加"
	DEL    = "删除"
	UPDATE = "修改"
	QUERY  = "查询"
	OTHER  = "其它"
)

func LoginAndOperationLogToDB(c *gin.Context) {
	// 日志
	var ol = new(models.OperationLog)

	method := c.Request.Method
	uri := c.Request.RequestURI
	ip := c.ClientIP()

	// 赋值
	ol.Method = method
	ol.IP = ip
	ol.URI = uri
	ol.CreateTime = time.Now()

	// 判断操作模块
	if strings.Contains(uri, "/login") {
		ol.ModuleName = "用户登录"
		ol.OperationType = LOGIN
	} else if strings.Contains(uri, "/logout") {
		ol.ModuleName = "用户登出"
		ol.OperationType = LOGOUT
	} else if strings.Contains(strings.ToLower(uri), "role") {
		ol.ModuleName = "角色管理"
		ol.OperationType = ObtainOperationType(uri)
	} else if strings.Contains(strings.ToLower(uri), "user") {
		ol.ModuleName = "用户管理"
		ol.OperationType = ObtainOperationType(uri)
	} else if strings.Contains(strings.ToLower(uri), "release") {
		ol.ModuleName = "发行管理"
		ol.OperationType = ObtainOperationType(uri)
	} else if strings.Contains(strings.ToLower(uri), "operation") {
		ol.ModuleName = "操作日志"
		ol.OperationType = ObtainOperationType(uri)
	} else if strings.Contains(strings.ToLower(uri), "loggin") {
		ol.ModuleName = "登录日志"
		ol.OperationType = ObtainOperationType(uri)
	} else {
		ol.ModuleName = "其它"
		ol.OperationType = "其它"
	}

	// 执行业务代码
	c.Next()

	// 获取执行状态
	userData, ok := c.Get("current_user")
	currentUser := userData.(models.User)
	// 登录成功
	if ok {
		ol.UserId = currentUser.UserId
		code := c.Writer.Status()
		if code == 200 {
			ol.Status = "成功"
		} else {
			ol.Status = "失败"
		}

		// 登录失败
	} else {
		var user models.User
		c.ShouldBind(&user)
		m := new(models.User)
		m.FindUserByUsername(user.Username)

		ol.UserId = m.UserId
		ol.Status = "失败"
	}

	// 存入数据库
	ol.AddOperationLog(ol)
}

// 判断操作类型
func ObtainOperationType(uri string) string {
	if strings.Contains(strings.ToLower(uri), "add") {
		return ADD
	} else if strings.Contains(strings.ToLower(uri), "delete") {
		return DEL
	} else if strings.Contains(strings.ToLower(uri), "update") {
		return UPDATE
	} else if strings.Contains(strings.ToLower(uri), "query") {
		return QUERY
	} else {
		return OTHER
	}
}
