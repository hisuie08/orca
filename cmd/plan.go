/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/ /*
未実装
*/
package cmd

import (
	"io"
	"orca/internal/context"
	"orca/process"
	"os"

	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "orcaがcomposeを管理する際変更される箇所を出力",

	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		return runPlan(cwd, os.Stdout)
	},
}

func runPlan(orcaRoot string, w io.Writer) error {
	// TODO: mode implement
	ctx, err := context.BuildContext(orcaRoot, w, context.ModeExecute)
	if err != nil {
		return err
	}
	process.PlanProcess(*ctx)
	// o, _ := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// printer.W = o
	// printer.C.Enabled = false

	return nil
}

func init() {
	rootCmd.AddCommand(planCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
