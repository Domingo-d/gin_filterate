/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/30 11:59
 */

package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server/config"
	"server/model"
)

var (
	AhoCorasick *model.AhoCorasick
	Router      *gin.Engine
	Redis       redis.UniversalClient
	Logger      *zap.Logger

	VP     *viper.Viper
	Config *config.ServerConfig
)
