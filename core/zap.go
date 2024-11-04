package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"server/global"
	"server/utils"
	"time"
)

func Zap() *zap.Logger {
	if ok, _ := utils.PathExists(global.Config.ZapConf.Direct); !ok {
		fmt.Printf("create %v directory \n", global.Config.ZapConf.Direct)

		_ = os.Mkdir(global.Config.ZapConf.Direct, os.ModePerm)
	}

	cores := GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if global.Config.ZapConf.LogInConsole {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.Config.ZapConf.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		path := fmt.Sprintf("%s/%s.log", global.Config.ZapConf.Direct, level.String())
		var writer = zapcore.AddSync(&lumberjack.Logger{
			Filename:   path, // 日志文件的位置
			MaxSize:    10,   // 文件最大尺寸（以MB为单位）
			MaxBackups: 3,    // 保留的最大旧文件数量
			MaxAge:     28,   // 保留旧文件的最大天数
			Compress:   true, // 是否压缩/归档旧文件
			LocalTime:  true, // 使用本地时间创建时间戳
		})

		cores = append(cores, zapcore.NewCore(GetEncoder(), writer, level))
	}

	if global.Config.ZapConf.LogInConsole {
		cores = append(cores, zapcore.NewCore(GetEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	}

	return cores
}

func GetEncoder() zapcore.Encoder {
	if global.Config.ZapConf.Format == "json" {
		return zapcore.NewJSONEncoder(GetEncoderConfig())
	}

	return zapcore.NewConsoleEncoder(GetEncoderConfig())
}

func GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "tims",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: global.Config.ZapConf.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		//EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeTime:   getEncodeTime,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}

func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.DateTime))
}
