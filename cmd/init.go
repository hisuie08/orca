/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"orca/internal/config"
	"orca/internal/ostools"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create orca.yml",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		current, _ := os.Getwd()
		orcaFile := current + "/orca.yml"
		if !ostools.FileExisists(orcaFile) {
			fmt.Printf("creating orca.yml\n")
			config.Create(orcaFile)
			fmt.Printf("%v created.\nyou can launch orca\n", orcaFile)
		}
	},
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
