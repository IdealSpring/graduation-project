package router

import (
	"github.com/gin-gonic/gin"
	. "graduation-project/server/controller"
	"graduation-project/server/middleware"
)

// 初始化路由
func InitRouter()*gin.Engine {
	r := gin.Default()

	// 跨域处中间件
	r.Use(middleware.CrossDomain)
	// 统一异常处理中间件
	r.Use(middleware.CustomError)

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
		auth.POST("/roleQuery", PostRoleQuery)
	}

	// 发行省份相关路由
	release := r.Group("/auth/release")
	release.Use(middleware.AuthMiddleware)
	{
		release.POST("/queryProvince", PostQueryProvince)
	}

	return r
}
