/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 16:12
 */

package router

import (
	"filterate/api"
	"github.com/gin-gonic/gin"
)

type (
	FilterateApi struct{}
)

func (receiver *FilterateApi) InitFilterateRouter(router *gin.RouterGroup) {
	apiRouter := router.Group("api")
	apiRouterApi := api.ApiGroupApp.FilterateApi

	{
		apiRouter.POST("filterate", apiRouterApi.Filterate)
	}
}