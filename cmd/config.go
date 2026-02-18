/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"orca/internal/capability"
	"orca/internal/logger"
	"orca/internal/usecase/config"
	"orca/model/policy/log"
	. "orca/presenter/formatter/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
// TODO: Process層に分離
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Reads and displays the current orca.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		caps := capability.BuildCommandCaps(*cmd)
		out := logger.New(caps)
		cfg, err := config.Load(caps)
		if err != nil {
			return err
		}
		out.Logln(log.LogNormal, FmtConfig(*cfg))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
