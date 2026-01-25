package context

import (
	"io"
	"orca/cmd/option"

	"github.com/spf13/cobra"
)

type CommandContext interface {
	WithRoot
	WithLog
}

func BuildBaseContext(cmd cobra.Command, opt option.BaseOption) CommandContext {
	ctx := New().
		WithRoot(opt.Root)
	return &ctx
}

func FromCommandCtx(ctx CommandContext) Context {
	c := New().WithRoot(ctx.Root()).WithLog(ctx.LogLevel(), ctx.LogTarget())
	return c
}

type silentWriter struct{}

func (s *silentWriter) Write([]byte) (int, error) {
	return 0, nil
}

func GetWriter(out io.Writer, disablef ...bool) io.Writer {
	for _, b := range disablef {
		if b {
			return &silentWriter{}
		}
	}
	return out
}
