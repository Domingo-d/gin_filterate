package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type (
	ZapConf struct {
		Level         string `mapstructure:"level" json:"level" yaml:"level"`                         // 日志级别
		Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                      // 日志前缀
		Format        string `mapstructure:"format" json:"format" yaml:"format"`                      // 日志输出格式
		Direct        string `mapstructure:"direct" json:"direct" yaml:"direct"`                      // 日志输出路径
		EncodeLevel   string `mapstructure:"encodeLevel" json:"encodeLevel" yaml:"encodeLevel"`       // 编码器配置
		StacktraceKey string `mapstructure:"stacktraceKey" json:"stacktraceKey" yaml:"stacktraceKey"` // 栈追踪

		MaxAge       int  `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`                   // 日志文件最大保存时间
		ShowLine     bool `mapstructure:"showLine" json:"showLine" yaml:"showLine"`             // 显示行号
		LogInConsole bool `mapstructure:"logInConsole" json:"logInConsole" yaml:"logInConsole"` // 输出到控制台
	}
)

// TransportLevel 根据字符串转化为 zapcore.Level
func (z *ZapConf) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
