package core

import (
	"GinWell-Server/global"
	"GinWell-Server/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GW_CONFIG.Zap.Director); !ok {
		fmt.Printf("create #{global.GW_CONFIG.Zap.Director} directory\n")
		_ = os.Mkdir(global.GW_CONFIG.Zap.Director, os.ModePerm)
	}
	debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.DebugLevel
	})

	infoPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.InfoLevel
	})
	warnPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel
	})
	errorPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})
	now := time.Now().Format("2006-01-02")
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%s/debug.log", now), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/info.log", now), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/warn.log", now), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/%s/error.log", now), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.GW_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GW_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.GW_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GW_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GW_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GW_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getEncoder() zapcore.Encoder {
	if global.GW_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := utils.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GW_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
