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
	OrcaRoot         string
	Config           *config.ResolvedConfig
	RunMode          RunMode
	ComposeInspector *inspector.DockerComposeInspector
	DockerInspector  *inspector.DockerInspector
	Printer          *orca.Printer
}

func BuildContext(orcaRoot string, w io.Writer) (*OrcaContext, error) {
	result := &OrcaContext{
		OrcaRoot:         orcaRoot,
		ComposeInspector: &inspector.DockerComposeInspector{OrcaRoot: orcaRoot},
		DockerInspector:  &inspector.DockerInspector{},
		Printer:          orca.NewPrinter(w, *orca.NewColorizer(w)),
	}
	cfg, err := config.LoadConfig(orcaRoot, inspector.ConfigFileReader{})
	if err != nil {
		return nil, err
	}
	result.Config = cfg
	return result, nil
}
