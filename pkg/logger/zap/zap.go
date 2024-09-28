package zaplogger

import (
	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapLogger struct {
	logger *zap.Logger
}

func New() *ZapLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	zapLogger, _ := config.Build()
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Debug("[ ZAP ]", err.Error())
			return
		}
	}(zapLogger)
	return &ZapLogger{
		logger: zapLogger,
	}
}

func (l *ZapLogger) Info(prefix, msg string) {
	l.logger.Info(prefix,
		zap.String(helpers.INFO_PREFIX, msg),
	)
}

func (l *ZapLogger) Debug(prefix, msg string) {
	l.logger.Debug(prefix,
		zap.String(helpers.INFO_PREFIX, msg),
	)
}

func (l *ZapLogger) Fatal(prefix, msg string) {
	l.logger.Fatal(prefix,
		zap.String(helpers.INFO_PREFIX, msg),
	)
}

func (l *ZapLogger) Response(prefix, status, msg string) {
	l.logger.Info(prefix,
		zap.String(helpers.STATUS_PREFIX, status),
		zap.String(helpers.INFO_PREFIX, msg),
	)
}
