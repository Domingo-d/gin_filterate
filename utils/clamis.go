package utils

import (
	"filterate/global"
	"filterate/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}

	return token
}

func GenerateToken(username string) string {
	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime},
		UserName:       username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(global.Config.JwtKey))
	if err != nil {
		global.Logger.Error("************** 生成 jwt 失败", zap.Error(err))
		return ""
	}

	return signedToken
}

func ParseToken(tokenClaims string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenClaims, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JwtKey), nil
	})

	if !token.Valid {
		return nil, err
	}

	return token.Claims.(*model.Claims), err
}
