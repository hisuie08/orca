package option

import (
	"orca/cmd/baseflag"
	"orca/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

type BaseOption struct {
	LogLevel logger.LogLevel
	Root     string
}

func NewBaseOption(cmd cobra.Command) BaseOption {
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

	return BaseOption{Root: wd, LogLevel: logLevel}
}
