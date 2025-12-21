/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	orca "orca/helper"
	"orca/internal/config"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [cluster-name]",
	Short: "Initialize an orca cluster in current directory",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		clusterName := ""
		if len(args) == 1 {
			clusterName = args[0]
		}

		return runInit(cwd, clusterName)
	},
}

func runInit(baseDir, clusterName string) error {
	path := filepath.Join(baseDir, orca.OrcaYamlFile)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("orca.yml already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	cfg := config.Create(clusterName)
	if err := writeConfig(path, cfg); err != nil {
		return err
	}
	fmt.Printf("%v was created successfully\n", path)
	return nil
}

func writeConfig(path string, cfg *config.OrcaConfig) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
