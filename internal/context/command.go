package context

import (
	"orca/cmd/baseflag"
	"orca/model/policy"
	"orca/model/policy/log"
	"os"

	"github.com/spf13/cobra"
)

type CommandContext interface {
	WithRoot
	WithLog
	WithPolicy
}

func BuildCommandCtx(cmd cobra.Command) CommandContext {
	wd, err := os.Getwd()
	if err != nil {
		panic("can't get working directory")
	}
	silent, _ := cmd.Flags().GetBool(baseflag.Silent)
	debug, _ := cmd.Flags().GetBool(baseflag.Debug)
	dry, _ := cmd.Flags().GetBool(baseflag.DryRun)
	logLevel := func() log.LogLevel {
		if silent {
			return log.LogSilent
		} else if debug {
			return log.LogDetail
		} else {
			return log.LogNormal
		}
	}()
	plcy := func() policy.ExecPolicy {
		if dry {
			return policy.Dry
		}
		return policy.Real
	}()
	ctx := New().WithRoot(wd).WithPolicy(plcy).
		WithLog(logLevel, cmd.OutOrStdout())
	return &ctx
}

func FromCommandCtx(ctx CommandContext) Context {
	c := New().WithRoot(ctx.Root()).WithPolicy(ctx.Policy()).
		WithLog(ctx.LogLevel(), ctx.LogTarget())
	return c
}
