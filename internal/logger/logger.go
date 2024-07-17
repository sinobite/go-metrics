package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Logger struct {
	ZapLog *zap.SugaredLogger
}

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
}

func New(logLevel string) Logger {
	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logCfg.Level.SetLevel(logLevelMap[logLevel])

	logger, err := logCfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	return Logger{ZapLog: logger.Sugar()}
}
