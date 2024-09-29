/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:09
 */

package initialize

import (
	"filterate/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemGroup := Router.Group("")
	router.RouterGroupApp.InitFilterateRouter(systemGroup)

	return Router
}