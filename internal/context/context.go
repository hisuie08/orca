package context

import (
	"io"
	orca "orca/helper"
	"orca/infra/inspector"
	"orca/internal/config"
)

type RunMode int

const (
	ModeExecute RunMode = iota
	ModeDryRun
)

type PlanOption struct {
	Mode   RunMode
	Silent bool
}

type OrcaContext struct {
	OrcaRoot string
	Config   *config.ResolvedConfig
	RunMode  RunMode
	Printer  *orca.Printer
}

func BuildContext(orcaRoot string, w io.Writer, m RunMode) (*OrcaContext, error) {
	result := &OrcaContext{
		OrcaRoot: orcaRoot,
		RunMode:  m,
		Printer:  orca.NewPrinter(w, *orca.NewColorizer(w)),
	}
	cfg, err := config.LoadConfig(inspector.ConfigFile(orcaRoot))
	if err != nil {
		return nil, err
	}
	result.Config = cfg
	return result, nil
}
