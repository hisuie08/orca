package option

import (
	"orca/cmd/baseflag"
	"os"

	"github.com/spf13/cobra"
)

type BaseOption struct {
	Silent bool
	Debug  bool
	Root   string
}

func NewBaseOption(cmd cobra.Command) (BaseOption, error) {
	wd, err := os.Getwd()
	if err != nil {
		return BaseOption{}, err
	}
	silent, err := cmd.Flags().GetBool(baseflag.Silent)
	if err != nil {
		return BaseOption{}, err
	}
	debug, err := cmd.Flags().GetBool(baseflag.Debug)
	if err != nil {
		return BaseOption{}, err
	}

	return BaseOption{Root: wd, Silent: silent, Debug: debug}, nil
}
