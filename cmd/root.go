package cmd

import (
	"orca/cmd/baseflag"
	"orca/cmd/internal"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "orca",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	err := RootCmd.Execute()
	if err == nil {
		os.Exit(0)
	}
	silent := func() bool {
		if s, e := RootCmd.Flags().GetBool(baseflag.Silent); e != nil {
			return false
		} else {
			return s
		}
	}()
	exitCode := internal.HandleError(err, silent)
	os.Exit(exitCode)
}

func init() {
	RootCmd.PersistentFlags().Bool(baseflag.DryRun, false, "Run in DRY mode")
	RootCmd.PersistentFlags().Bool(baseflag.Debug, false, "Show detail logs")
	RootCmd.PersistentFlags().Bool(baseflag.Silent, false, "Never print to console")
}
