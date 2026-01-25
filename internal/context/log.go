package context

import (
	"io"
	"orca/internal/logger"
)

type WithLog interface {
	LogLevel() logger.LogLevel
	LogTarget() io.Writer
}

type withLog struct {
	logLevel logger.LogLevel
	out      io.Writer
}

func (l *withLog) LogLevel() logger.LogLevel {
	return l.logLevel
}

func (l *withLog) LogTarget() io.Writer {
	return l.out
}
