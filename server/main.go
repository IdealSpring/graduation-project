package main

import (
	"github.com/gin-gonic/gin"
	lg "graduation-project/server/common/logger"
	"graduation-project/server/config"
	orm "graduation-project/server/database"
	"graduation-project/server/router"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// 初始化系统数据库
	orm.InitApplicationDatabase()

	// 初始化路由
	r := router.InitRouter()

	// 关闭数据库
	defer orm.DB.Close()

	// 启动服务
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		lg.Logln(err)
	}
}
