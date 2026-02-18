package plan

import (
	"orca/errs"
	"orca/internal/capability"
	"orca/internal/executor"
	"orca/internal/logger"
	"orca/internal/usecase/compose"
	"orca/internal/usecase/config"
	"orca/internal/usecase/dotorca"
	"orca/internal/usecase/plan"

	"gopkg.in/yaml.v3"
)

type PlanOption struct {
	Force bool
}
type planProcessCapability interface {
	capability.CommandCapabilities
	capability.WithConfig
	capability.WithColor
}

type Process struct {
	caps   capability.CommandCapabilities
	logger logger.Logger
}

func New(c capability.CommandCapabilities) *Process {
	return &Process{caps: c, logger: logger.New(c)}
}

func (p *Process) Run(o PlanOption) error {
	cfg, err := config.Load(p.caps)
	if err != nil {
		return err
	}
	caps := func() planProcessCapability {
		c := capability.FromCommandCaps(p.caps).WithConfig(cfg).
			WithColor(p.caps.LogTarget())
		return &c
	}()
	return p.run(caps, o)
}

func (p *Process) run(caps planProcessCapability, o PlanOption) error {
	cmp, err := compose.GetAllCompose(caps)
	if err != nil {
		return errs.ErrComposeNotFound
	}
	orcaplan := plan.BuildOrcaPlan(caps, cmp)
	overlayer := compose.ComposeOverlayer(caps, cmp)
	if caps.Config().Volume.Enabled() {
		overlayer.OverlayVolume(orcaplan.Volumes)
	}
	if caps.Config().Network.Enabled {
		overlayer.OverlayNetwork(orcaplan.Networks)
	}
	dotorca.DumpComposes(caps, cmp, o.Force)
	pl, _ := yaml.Marshal(orcaplan)
	executor.NewFilesystem(p.caps).WriteFile(p.caps.OrcaPlanFile(), pl)
	return nil
}
