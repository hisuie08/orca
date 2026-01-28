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
	silent, e := RootCmd.Flags().GetBool(baseflag.Silent)
	if e != nil {
		silent = false
	}
	exitCode := internal.HandleError(e, silent)
	os.Exit(exitCode)
}

func init() {
	RootCmd.PersistentFlags().Bool(baseflag.DryRun, false, "Run in DRY mode")
	RootCmd.PersistentFlags().Bool(baseflag.Debug, false, "Show detail logs")
	RootCmd.PersistentFlags().Bool(baseflag.Silent, false, "Never print to console")
}
