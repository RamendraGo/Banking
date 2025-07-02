package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	//Override log config
	config := zap.NewProductionConfig()
	encoderconfig := zap.NewProductionEncoderConfig()

	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderconfig.TimeKey = "timestamp"
	encoderconfig.StacktraceKey = ""
	config.EncoderConfig = encoderconfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {

	log.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {

	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {

	log.Error(message, fields...)
}
