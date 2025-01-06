package zapLogger

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})
	Panic(...interface{})
	Debug(...interface{})
}

func InitZap() *zap.SugaredLogger {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	zapConfig.EncoderConfig.TimeKey = "time"
	l, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("ZapConfig loading error %v", err.Error())
	}

	logger := l.Sugar()

	return logger
}
