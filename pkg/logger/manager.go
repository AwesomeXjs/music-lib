package logger

type Logger interface {
	Info(prefix, msg string)
	Fatal(prefix, msg string)
	Response(prefix, status, msg string)
}
