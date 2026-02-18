package capability

import (
	"orca/cmd/baseflag"
	"orca/model/policy"
	"orca/model/policy/log"
	"os"

	"github.com/spf13/cobra"
)

type CommandCapabilities interface {
	WithRoot
	WithLog
	WithPolicy
}

func BuildCommandCaps(cmd cobra.Command) CommandCapabilities {
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
	caps := New().WithRoot(wd).WithPolicy(plcy).
		WithLog(logLevel, cmd.OutOrStdout())
	return &caps
}

func FromCommandCaps(caps CommandCapabilities) Capability {
	c := New().WithRoot(caps.Root()).WithPolicy(caps.Policy()).
		WithLog(caps.LogLevel(), caps.LogTarget())
	return c
}
