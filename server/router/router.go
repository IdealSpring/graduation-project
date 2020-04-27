package router

import (
	"github.com/gin-gonic/gin"
	. "graduation-project/server/controller"
	"graduation-project/server/middleware"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	// 跨域处中间件
	r.Use(middleware.CrossDomain)
	// 统一异常处理中间件
	r.Use(middleware.CustomError)
	// 登录日志和操作日志
	r.Use(middleware.LoginAndOperationLogToDB)

	// 添加路由规则
	r.POST("/login", LoginHandler)

	// 用户相关路由
	auth := r.Group("/auth/user")
	auth.Use(middleware.AuthMiddleware)
	{
		// 登录相关
		auth.GET("/info", GetUserInfo)

		// 右上角设置相关
		auth.PUT("/pwd", PutUserPwd)
		auth.POST("/logout", PostLogout)

		// 用户管理相关
		auth.GET("/option/role", GetRole)
		auth.POST("/query", PostQuery)
		auth.DELETE("/delete/:uid", DeleteUserById)
		auth.POST("/addUser", PostAddUser)
		auth.PUT("/update", PutUpdateUser)
		auth.PUT("/updateRole", UpdateRole)

		// 角色管理相关
		auth.POST("/roleQuery", PostRoleQueryToPage)
		auth.GET("/rolePerms/:roleId", GetRolePerms)
		auth.PUT("/role/updatePerm", PutUpdatePerm)
		auth.POST("/addRole", PostAddRole)
		auth.PUT("/role/updateRole", PutUpdateRole)
		auth.DELETE("/role/delete", DeleteRole)
	}

	// 发行省份相关路由
	release := r.Group("/auth/release")
	release.Use(middleware.AuthMiddleware)
	{
		release.POST("/queryProvince", PostQueryProvince)
		release.GET("/option/province", GetAllProvice)
		release.POST("/addProvince", PostAddProvince)
	}

	// 日志相关路由
	operationLog := r.Group("/auth/log")
	operationLog.Use(middleware.AuthMiddleware)
	{
		// 登录日志
		operationLog.DELETE("/loggin/deleteAll", DeleteDeleteAllLoginLog)

		// 操作日志
		operationLog.POST("/operation/queryDataToPage", PostFetchDataToPage)
		operationLog.DELETE("/operation/delete/:logId", DeleteOperationLog)
		operationLog.DELETE("/operation/deleteAll", DeleteDeleteAllOperationLog)
	}

	return r
}
