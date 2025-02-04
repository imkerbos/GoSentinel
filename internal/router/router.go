/**
 * @Auth: Kerbos - 0xkerbos@gmail.com
 * @Blog: https://www.gov-cn.cn
 * @Date: 9/11/2023 - 11:40 pm
 * @File: internal/router/router.go
 * @Desc:
 */

package router

import (
	"GoSentinel/internal/controller"
	"GoSentinel/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 全局中间件
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())

	// JWT中间件
	//authMiddleware := middleware.JwtMiddleware()

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", middleware.CheckJWTAndRespond(), controller.IndexPage)
	router.GET("/login", controller.LoginPage)

	// 登录认证路由组
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controller.LoginHandler)
		authGroup.GET("/captcha", controller.GenerateCaptcha)
	}

	// API路由组
	//apiGroup := router.Group("/api")
	//apiGroup.Use(authMiddleware)
	//{
	//	//apiGroup.GET("/service", controller.GetService)
	//	//apiGroup.POST("/service", controller.CreateService)
	//	// 更多API路由...
	//}

	// 监控路由组
	//monitorGroup := router.Group("/monitor")
	//{
	//	monitorGroup.GET("/metrics", controller.GetMetrics)
	//	// 更多监控路由...
	//}

	return router
}
