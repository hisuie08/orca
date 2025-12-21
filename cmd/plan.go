/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/ /*
未実装
*/
package cmd

import (
	"fmt"
	"io"
	orca "orca/helper"
	"orca/internal/compose"
	"orca/internal/config"
	"orca/internal/plan"
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

func newCfgReader(p string)*config.ConfigFileReader{
	return  &config.ConfigFileReader{OrcaRoot: p}
}
func runPlan(orcaRoot string, w io.Writer) error {
	printer := orca.NewPrinter(w, *orca.NewColorizer(w))
	cfg, err := config.LoadConfig(orcaRoot,newCfgReader(orcaRoot))
	if err != nil {
		return err
	}
	cmp, err := compose.GetAllCompose(orcaRoot,compose.DockerComposeInspector{})
	if err != nil {
		return err
	}
	vol := cmp.CollectVolumes()
	net := cmp.CollectComposes()
	volumePlan := plan.BuildVolumePlan(vol, cfg.Volume)

	networkPlan := plan.BuildNetworkPlan(net, cfg.Network)
	plan.PrintVolumePlanTable(volumePlan, printer)

	fmt.Printf("\n")

	plan.PrintNetworkPlan(networkPlan, printer)
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
