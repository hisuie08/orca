package cmd

import (
	"orca/internal/context"
	planprocess "orca/internal/process/plan"

	"github.com/spf13/cobra"
)

func newPlanCommand() *cobra.Command {
	var opt planprocess.PlanOption = planprocess.PlanOption{}
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "orcaがcomposeを管理する際変更される箇所を出力",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.BuildCommandCtx(*cmd)
			proc := planprocess.New(ctx)
			return proc.Run(opt)
		},
	}
	return cmd
}

func init() {
	RootCmd.AddCommand(newPlanCommand())
}
