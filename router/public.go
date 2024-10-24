package router

import (
	"filterate/api"
	"github.com/gin-gonic/gin"
)

type (
	PublicApi struct{}
)

func (receiver *PublicApi) InitPublicRouter(router *gin.RouterGroup) {
	loginGroup := router.Group("/login")
	loginApi := api.ApiGroupApp.LoginApi

	{
		loginGroup.POST("/signIn", loginApi.SignIn)
		loginGroup.POST("/signUp", loginApi.SignUp)
		loginGroup.POST("/signOut", loginApi.SignOut)
	}
}
