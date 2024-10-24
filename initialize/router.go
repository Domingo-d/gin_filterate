/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:09
 */

package initialize

import (
	"filterate/global"
	"filterate/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapio"
	"net/http"
)

func Routers() *gin.Engine {
	gin.DefaultWriter = &zapio.Writer{Log: global.Logger}
	gin.DefaultErrorWriter = gin.DefaultWriter

	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	//Router.Use(gin.LoggerWithWriter(gin.DefaultWriter))

	global.Logger.Info("替换 gin Logger 成功")
	global.Logger.Info("开始注册路由")

	publicGroup := Router.Group("")
	{
		publicGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	privateGroup := Router.Group("")
	privateGroup.Use()
	router.InitRouter(publicGroup, privateGroup)

	global.Logger.Info("路由注册完成")

	return Router
}
