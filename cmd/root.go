/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"orca/cmd/cmdflag"
	"orca/errs"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "orca",
	Short: "",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err == nil {
		os.Exit(0)
	}

	exitCode := handleError(*rootCmd, err)
	os.Exit(exitCode)
}
func handleError(cmd cobra.Command, err error) int {
	silent, err := cmd.PersistentFlags().GetBool(cmdflag.Silent)
	if err != nil {
		silent = false
	}
	switch {
	case errors.Is(err, errs.ErrAlreadyInitialized):
		if !silent {
			fmt.Fprintln(os.Stderr, "このディレクトリは既に初期化されています")
			fmt.Fprintln(os.Stderr, "--force オプションで再生成してください")
		}
		return 0
	case errors.Is(err, errs.ErrNotInitialized):
		if !silent {
			fmt.Fprintln(os.Stderr, "このディレクトリは初期化されていません")
			fmt.Fprintln(os.Stderr, "orca init を先に実行してください")
		}
		//return exitcode.NotInitialized
		return 1

	case errors.Is(err, errs.ErrPlanDirty):
		if !silent {
			fmt.Fprintln(os.Stderr, "plan が最新ではありません")
			fmt.Fprintln(os.Stderr, "orca plan を再実行してください")
		}
		//return exitcode.PlanDirty
		return 1

	case errors.Is(err, errs.ErrDryRunViolation):
		if !silent {
			fmt.Fprintln(os.Stderr, "dry-run のため操作は実行されませんでした")
		}
		//return exitcode.OK
		return 0

	default:
		if !silent {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		//return exitcode.GeneralError
		return 1
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orca.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().Bool(cmdflag.DryRun, false, "Run in DRY mode")
	rootCmd.PersistentFlags().Bool(cmdflag.Debug, false, "Show detail logs")
	rootCmd.PersistentFlags().Bool(cmdflag.Silent, false, "Never print to console")
}
