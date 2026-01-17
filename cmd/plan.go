/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/ /*
未実装
*/
package cmd

import (
	"orca/internal/context"
	"orca/model/policy"
	"orca/process"
	"orca/process/plan"
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
		ctx := context.New().WithRoot(cwd).
			WithPolicy(policy.Real).WithReport(cmd.OutOrStderr())
		return runPlan(&ctx)
	},
}

type PlanContext interface {
	context.WithRoot
	context.WithPolicy
	context.WithReport
}

func runPlan(ctx PlanContext) error {
	// TODO: mode implement
	// o, _ := os.OpenFile("./log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// printer.W = o
	// printer.C.Enabled = false
	process.PlanProcess(ctx, plan.PlanOption{Force: true})
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
