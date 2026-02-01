package logger

import (
	"bytes"
	"orca/internal/context"
	. "orca/model/policy/log"
	"testing"
)

func TestCompareLogLevel(t *testing.T) {
	tests := []struct {
		desc    string
		lp      LogLevel
		ll      LogLevel
		willLog bool
	}{
		{desc: "Silent-Normal", lp: LogSilent, ll: LogNormal, willLog: false},
		{desc: "Normal-Debug", lp: LogNormal, ll: LogDetail, willLog: false},
		{desc: "Debug-Normal", lp: LogDetail, ll: LogNormal, willLog: true},
		{desc: "Normal-Normal", lp: LogNormal, ll: LogNormal, willLog: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx := context.New().WithLog(tt.lp, new(bytes.Buffer))
			lg := New(&ctx)
			l := lg.chkPolicy(tt.ll)
			if tt.willLog != l {
				t.Errorf("expected %t but got %t", tt.willLog, l)
			}
		})
	}
}
