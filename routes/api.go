package routes

import (
	"go-api/api/app/controllers/admin"
	"go-api/api/app/middleware"

	"github.com/gin-gonic/gin"
)

// 注册路由列表
func ApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/admin/login", admin.Login) // 登录

	// 登录鉴权路由
	auth := router.Group("api")             // 认证路由组
	auth.Use(middleware.ApiAuth())          // 登录认证中间件
	auth.GET("/admin/auth", admin.AuthInfo) // 登录用户信息

	// start admin
	api.POST("/admin/create", admin.Create)
	api.POST("/admin/update", admin.Update)
	api.GET("/admin/delete", admin.Delete)
	api.GET("/admin/info", admin.Info)
	api.POST("/admin/paginate", admin.Paginate)
	// end admin
}
