package logger

import "log"

type stdoutLogger struct{}

func (l *stdoutLogger) Write(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}

func NewStdoutLogger() Logger {
	return &stdoutLogger{}
}
