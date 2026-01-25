package context

import (
	"io"
	"orca/cmd/baseflag"
	"orca/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

type CommandContext interface {
	WithRoot
	WithLog
}

func BuildCommandCtx(cmd cobra.Command) CommandContext {
	wd, err := os.Getwd()
	if err != nil {
		panic("can't get working directory")
	}
	silent, _ := cmd.Flags().GetBool(baseflag.Silent)
	debug, _ := cmd.Flags().GetBool(baseflag.Debug)
	logLevel := func() logger.LogLevel {
		if silent {
			return logger.LogSilent
		} else if debug {
			return logger.LogDebug
		} else {
			return logger.LogNormal
		}
	}()
	ctx := New().WithRoot(wd).WithLog(logLevel, cmd.OutOrStdout())
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
