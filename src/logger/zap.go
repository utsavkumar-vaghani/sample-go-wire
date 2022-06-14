package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		var cfg zap.Config
		if os.Getenv("ENV") != "" || os.Getenv("ENV") != "local" {
			cfg = zap.NewProductionConfig()
			cfg.OutputPaths = []string{"stdout"}
			cfg.ErrorOutputPaths = []string{"stdout"}
			cfg.InitialFields = map[string]interface{}{"name": "sample-go-wire"}
			// cfg.EncoderConfig.EncodeLevel = IntegerLevelEncoder
			cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			cfg.EncoderConfig.TimeKey = "time"

		} else {
			cfg = zap.NewDevelopmentConfig()
		}
		l, err := cfg.Build()
		if err != nil {
			panic(err)
		}

		logger = l.Sugar()
	})

	return logger
}
