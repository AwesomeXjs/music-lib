package logger

// Logger default interface for logger
type Logger interface {
	Info(prefix, msg string)
	Debug(prefix, msg string)
	Fatal(prefix, msg string)
	Response(prefix, status, msg string)
}
