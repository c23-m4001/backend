package infrastructure

import internalLogger "capstone/internal/logger"

type LoggerStack interface {
	WriteAll(format string, v ...interface{})
}

type loggerStack []internalLogger.Logger

func (i loggerStack) WriteAll(format string, v ...interface{}) {
	for _, logger := range i {
		logger.Write(format, v...)
	}
}

func NewLoggerStack(logChannels []string) LoggerStack {
	stack := loggerStack{}
	for _, logChannel := range logChannels {
		switch logChannel {
		case "stdout":
			stack = append(stack, internalLogger.NewStdoutLogger())
		}
	}

	return stack
}
