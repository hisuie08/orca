package context

import (
	"io"
	"orca/model/policy/log"
)

type WithLog interface {
	LogLevel() log.LogLevel
	LogTarget() io.Writer
}

type withLog struct {
	logLevel log.LogLevel
	out      io.Writer
}

func (l *withLog) LogLevel() log.LogLevel {
	return l.logLevel
}

func (l *withLog) LogTarget() io.Writer {
	return l.out
}
