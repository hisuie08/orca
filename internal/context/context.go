package context

import (
	"io"
	orca "orca/helper"
	"orca/infra/applier"
	ins "orca/infra/inspector"
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
type Inspectors struct {
	Compose *ins.DockerComposeInspector
	Docker  *ins.DockerInspector
}

type OrcaContext struct {
	OrcaRoot string
	Config   *config.ResolvedConfig
	RunMode  RunMode
	Printer  *orca.Printer
	Applier  applier.Appliers
}

func (o OrcaContext) NewInsCompose() *ins.DockerComposeInspector {
	return ins.NewInsCompose(o.OrcaRoot)
}

func (o OrcaContext) NewInsDocker() *ins.DockerInspector {
	return ins.NewInsDocker()
}

func BuildContext(orcaRoot string, w io.Writer, m RunMode) (*OrcaContext, error) {
	Applier := applier.Appliers{
		Compose: &applier.DotOrcaDumper{OrcaRoot: orcaRoot},
	}
	result := &OrcaContext{
		OrcaRoot: orcaRoot,
		Applier:  Applier,
		RunMode:  m,
		Printer:  orca.NewPrinter(w, *orca.NewColorizer(w)),
	}
	cfg, err := config.LoadConfig(orcaRoot, ins.ConfigFile{OrcaRoot: orcaRoot})
	if err != nil {
		return nil, err
	}
	result.Config = cfg
	return result, nil
}
