package logger

import (
	"io"
)

type LogLevel int

const (
	LogSilent LogLevel = iota
	LogNormal
	LogDebug
)

type logger interface {
	Log([]byte)
}

type Logger struct {
	out       io.Writer
	logPolicy LogLevel
	logLevel  LogLevel
}
func New(o io.Writer,lp LogLevel)Logger{
	return Logger{out:o, logPolicy: lp}
}
func (l Logger) Init(lv LogLevel) Logger {
	l.logLevel = lv
	return l
}
func (l *Logger) chkPolicy() bool {
	return l.logPolicy >= l.logLevel
}
func (l *Logger) Log(p []byte) {
	if l.chkPolicy() {
		l.out.Write(p)
	}
}
