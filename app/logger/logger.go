package logger

import (
	"github.com/natefinch/lumberjack"
	"go-rookie/app/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

func Init() (err error) {
	writeSyncer := getLogWriter(settings.Config.LogConfig.FileName,
		settings.Config.LogConfig.MaxSize,
		settings.Config.LogConfig.MaxBackups,
		settings.Config.LogConfig.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(settings.Config.LogConfig.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if settings.Config.Mode == "dev" {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
				zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
