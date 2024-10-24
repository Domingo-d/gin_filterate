package api

import (
	"filterate/model/request"
	"github.com/gin-gonic/gin"
)

type (
	LoginApi struct{}
)

func (l *LoginApi) SignIn(c *gin.Context) {
	info := &request.SignInReq{}
	if err := c.ShouldBindJSON(info); nil != err {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	code, token, err := userService.SignIn(info)
	if nil != err {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.JSON(code, gin.H{"msg": "token generated"})
}

func (l *LoginApi) SignUp(c *gin.Context) {
	info := &request.SignUpReq{}
	if err := c.ShouldBindJSON(info); nil != err {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	code, err := userService.SignUp(info)
	if nil != err {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"msg": "ok"})
}

func (l *LoginApi) SignOut(c *gin.Context) {

}
