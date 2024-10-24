/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:48
 */

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	FilterateRes struct {
		Str string `json:"str"`
	}

	Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
)

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: 7,
		Data: nil,
		Msg:  message,
	})
}
