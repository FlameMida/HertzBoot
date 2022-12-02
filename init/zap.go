package init

import (
	"HertzBoot/pkg"
	"HertzBoot/pkg/global"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var level zap.AtomicLevel

func Zap() {
	var logger *hertzZap.Logger
	if ok, _ := pkg.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	switch global.CONFIG.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	var zapOption = hertzZap.WithZapOptions()
	if global.CONFIG.Zap.ShowLine {
		zapOption = hertzZap.WithZapOptions(zap.AddCaller())
	}
	if level.Level() == zap.NewAtomicLevelAt(zap.DebugLevel).Level() || level.Level() == zap.NewAtomicLevelAt(zap.ErrorLevel).Level() {
		logger = hertzZap.NewLogger(
			hertzZap.WithCoreEnc(getEncoder()),
			hertzZap.WithCoreLevel(level),
			hertzZap.WithCoreWs(getEncoderCore()),
			hertzZap.WithZapOptions(
				zap.AddStacktrace(level),
			),
			zapOption,
		)
	} else {
		logger = hertzZap.NewLogger(
			hertzZap.WithCoreEnc(getEncoder()),
			hertzZap.WithCoreLevel(level),
			hertzZap.WithCoreWs(getEncoderCore()),
			zapOption,
		)
	}
	hlog.SetLogger(logger)
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "Msg",
		LevelKey:       "Level",
		TimeKey:        "Time",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() zapcore.WriteSyncer {
	writer, err := pkg.GetWriteSyncer() // 使用file-rotateLogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return writer
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
