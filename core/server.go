/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/30 11:54
 */

package core

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"net/http"
	"os"
	"os/signal"
	"server/global"
	"server/initialize"
	"syscall"
	"time"
)

func RunServer() {
	rand.Seed(uint64(time.Now().UnixNano()))

	confName := flag.String("confName", "srvConf.yaml", "配置文件名称")
	global.VP = Viper(*confName)

	global.Logger = Zap()
	defer global.Logger.Sync()

	gin.ForceConsoleColor()

	global.Logger.Info("------------------- zap初始化完毕 -------------------")

	zap.ReplaceGlobals(global.Logger)

	Redis()

	global.AhoCorasick = initialize.NewAhoCorasick()
	if nil == global.AhoCorasick.ReadPattern(global.Config.FilterateName) {
		global.AhoCorasick.Build()
	}

	global.Router = initialize.Routers()

	s := &http.Server{Addr: global.Config.System.Addr, Handler: global.Router}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		si := <-signals
		global.Logger.Info("接收到关闭信号", zap.Any("signal", si.String()))

		cancel()

		s.Shutdown(ctx)
	}()
	// 218*171*133mm
	global.Logger.Info("----------------------- server Start", zap.String("addr", global.Config.System.Addr))

	if global.Config.CertFile != "" && global.Config.KeyFile != "" {
		global.Logger.Info("-----------------------", zap.String("exit info",
			s.ListenAndServeTLS(global.Config.CertFile, global.Config.KeyFile).Error()))
	} else {
		global.Logger.Info("-----------------------", zap.String("exit info",
			s.ListenAndServe().Error()))
	}

	<-ctx.Done()

	global.Logger.Info("----------------------- server End")
}
