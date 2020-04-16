package log

import (
	"fmt"
	"github.com/figment-networks/oasis-rpc-proxy/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var (
	Log *zap.Logger
)

func init() {
	level, err := parseLevel(config.LogLevel())
	if err != nil {
		level = zap.DebugLevel
	}

	logConfig := zap.Config{
		OutputPaths: []string{"stdout", "/tmp/logs"},
		Encoding: "json",
		Level: zap.NewAtomicLevelAt(level),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func parseLevel(lvl string) (zapcore.Level, error) {
	switch strings.ToLower(lvl) {
	case "panic":
		return zap.PanicLevel, nil
	case "fatal":
		return zap.FatalLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	case "warn", "warning":
		return zap.WarnLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "debug":
		return zap.DebugLevel, nil
	}

	var l zapcore.Level
	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()
}

func Debug(msg string, tags ...zap.Field) {
	Log.Debug(msg, tags...)
	Log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.Error(msg, tags...)
	Log.Sync()
}
