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
}
func New(o io.Writer,lp LogLevel)Logger{
	return Logger{out:o, logPolicy: lp}
}
func (l *Logger) chkPolicy(lv LogLevel) bool {
	return l.logPolicy >= lv
}
func (l *Logger) Log(lv LogLevel,p []byte) {
	if l.chkPolicy(lv) {
		l.out.Write(p)
	}
}
