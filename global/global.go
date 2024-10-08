/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/30 11:59
 */

package global

import (
	"filterate/model"
	"github.com/gin-gonic/gin"
)

var (
	AhoCorasick *model.AhoCorasick
	Router      *gin.Engine
)