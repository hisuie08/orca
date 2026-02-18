package logger

import (
	"fmt"
	"io"
	"orca/internal/capability"
	. "orca/model/policy/log"
)

type logCapability interface {
	capability.WithLog
}
type logger interface {
	Log([]byte)
}

type Logger struct {
	out       io.Writer
	logPolicy LogLevel
}

func New(caps logCapability) Logger {
	return Logger{out: caps.LogTarget(), logPolicy: caps.LogLevel()}
}
func (l *Logger) chkPolicy(lv LogLevel) bool {
	return l.logPolicy >= lv
}
func (l *Logger) Log(lv LogLevel, p []byte) {
	if l.chkPolicy(lv) {
		l.out.Write(p)
	}
}

func (l *Logger) Logln(lv LogLevel, s string) {
	if l.chkPolicy(lv) {
		l.out.Write([]byte(fmt.Sprintf("%s\n", s)))
	}
}
