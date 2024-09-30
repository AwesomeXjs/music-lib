package zaplogger

import (
	"time"

	"github.com/AwesomeXjs/music-lib/internal/helpers"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger - struct for zap logger
type ZapLogger struct {
	logger *zap.Logger
}

// New - creates new zap logger
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

// Info - create info log
func (l *ZapLogger) Info(prefix, msg string) {
	l.logger.Info(prefix,
		zap.String(helpers.InfoPrefix, msg),
	)
}

// Debug - create debug log
func (l *ZapLogger) Debug(prefix, msg string) {
	l.logger.Debug(prefix,
		zap.String(helpers.InfoPrefix, msg),
	)
}

// Fatal - create fatal log
func (l *ZapLogger) Fatal(prefix, msg string) {
	l.logger.Fatal(prefix,
		zap.String(helpers.InfoPrefix, msg),
	)
}

// Response - create response log
func (l *ZapLogger) Response(prefix, status, msg string) {
	l.logger.Info(prefix,
		zap.String(helpers.StatusPrefix, status),
		zap.String(helpers.InfoPrefix, msg),
	)
}
