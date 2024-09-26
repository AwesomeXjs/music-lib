package zaplogger

import (
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func New() *ZapLogger {
	zapLogger, _ := zap.NewProduction()
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {

		}
	}(zapLogger)
	return &ZapLogger{
		logger: zapLogger,
	}
}

func (l *ZapLogger) Info(prefix, msg string) {
	l.logger.Info(prefix,
		zap.String(logger.INFO_PREFIX, msg),
	)
}

func (l *ZapLogger) Fatal(prefix, msg string) {
	l.logger.Fatal(prefix,
		zap.String(logger.INFO_PREFIX, msg),
	)
}

func (l *ZapLogger) Response(prefix, status, msg string) {
	l.logger.Info(prefix,
		zap.String(logger.STATUS_PREFIX, status),
		zap.String(logger.INFO_PREFIX, msg),
	)
}
