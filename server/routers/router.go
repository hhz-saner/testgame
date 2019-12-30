package routers

import (
	"fmt"
	"gameLog/config"
	"gameLog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	Router := gin.New()

	Router.Use(gin.Logger())

	Router.Use(gin.Recovery())

	Router.Use(Cors())

	if config.Cfg.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router(Router)
	return Router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func router(Router *gin.Engine) {
	Router.GET("/", controller.Ping)
	Router.POST("/login", controller.Login)
	Router.Any("/report", controller.Report)
}
