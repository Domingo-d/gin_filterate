/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:55
 */

package router

import (
	"filterate/middleware"
	"github.com/gin-gonic/gin"
)

type (
	RouterGroup struct {
		FilterateApi
		PublicApi
	}
)

var (
	RouterGroupApp = new(RouterGroup)
)

func InitRouter(public, private *gin.RouterGroup) {
	RouterGroupApp.InitPublicRouter(public)

	private.Use(middleware.JWTMiddleware())
	RouterGroupApp.InitFilterateRouter(private)
}
