package cmd

import (
	"orca/internal/context"
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
			// Process 呼び出し
			ctx := context.BuildCommandCtx(*cmd)
			proc := initprocess.New()
			return proc.Run(ctx, opt)
		},
	}

	// --- flags ---
	const (
		NoCreate        = "nocreate"
		Force           = "force"
		Volume          = "volume"
		EnsureVolume    = "ensure-volume"
		EnableNetwork   = "enable-network"
		NetworkName     = "network-name"
		NetworkInternal = "network-internal"
	)
	cmd.Flags().BoolVar(&opt.NoCreate, NoCreate, false, "Do not create orca.yml")
	cmd.Flags().BoolVar(&opt.Force, Force, false, "Overwrite existing orca.yml")
	// TODO: フラグで初期化オプションを指定可能にする
	// cmd.Flags().StringVar(&opt.Volume.Path, Volume, "", "Volume path")
	// cmd.Flags().BoolVar(&opt.Volume.EnsurePath, EnsureVolume, false, "Ensure volume path exists")
	// cmd.Flags().BoolVar(&opt.Network.Enabled, EnableNetwork, false, "Enable network")
	// cmd.Flags().StringVar(&opt.Network.Name, NetworkName, "", "Shared network name")
	// cmd.Flags().BoolVar(&opt.Network.Internal, NetworkInternal, false, "Make network internal")

	return cmd
}

func init() {
	RootCmd.AddCommand(newInitCommand())
}
