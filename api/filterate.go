/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/29 15:36
 */

package api

import (
	"github.com/gin-gonic/gin"
	"server/model/request"
)

type (
	FilterateApi struct{}
)

func (receiver *FilterateApi) Filterate(c *gin.Context) {
	info := &request.FilterateReq{}
	if err := c.ShouldBindJSON(info); nil != err {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	res := filterateService.Filter(info)

	c.JSON(200, res)
}
