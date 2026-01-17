/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	initprocess "orca/internal/process/init"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newInitCommand() *cobra.Command {
	var opt initprocess.InitOption

	cmd := &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize new orca cluster",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			if len(args) > 0 {
				opt.Name = args[0]
			} else {
				opt.Name = filepath.Base(wd)
			}
			ctx, err := BuildBaseContext(cmd)
			// Process 呼び出し
			proc := initprocess.New()
			return proc.Run(ctx, opt)
		},
	}

	// --- flags ---

	cmd.Flags().BoolVar(&opt.NoCreate, "nocreate", false, "Do not create orca.yml")
	cmd.Flags().BoolVar(&opt.Force, "force", false, "Overwrite existing orca.yml")

	cmd.Flags().StringVar(&opt.Volume.Path, "volume", "", "Volume path")
	cmd.Flags().BoolVar(&opt.Volume.EnsurePath, "ensure-volume-path", false, "Ensure volume path exists")

	cmd.Flags().BoolVar(&opt.Network.Enabled, "enable-network", false, "Enable network")
	cmd.Flags().StringVar(&opt.Network.Name, "network-name", "", "Shared network name")
	cmd.Flags().BoolVar(&opt.Network.Internal, "network-internal", false, "Make network internal")

	return cmd
}

func init() {
	rootCmd.AddCommand(newInitCommand())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
