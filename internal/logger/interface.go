package logger

type Logger interface {
	Write(format string, v ...interface{})
}
