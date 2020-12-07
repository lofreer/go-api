package bootstrap

import (
	"go-api/api/routes"

	"github.com/gin-gonic/gin" // 基于 gin 框架
)

func Start() *gin.Engine {
	// 数据库初始化
	autoMigrate()
	router := gin.Default()  // 获取路由实例
	routes.ApiRouter(router) // 注册路由
	return router            // 返回路由
}
