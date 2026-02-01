package plan

import (
	"orca/errs"
	"orca/internal/context"
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
type planProcessContext interface {
	context.CommandContext
	context.WithConfig
	context.WithColor
}

type Process struct {
	ctx    context.CommandContext
	logger logger.Logger
}

func New(c context.CommandContext) *Process {
	return &Process{ctx: c, logger: logger.New(c)}
}

func (p *Process) Run(o PlanOption) error {
	cfg, err := config.Load(p.ctx)
	if err != nil {
		return err
	}
	ctx := func() planProcessContext {
		c := context.FromCommandCtx(p.ctx).WithConfig(cfg).
			WithColor(p.ctx.LogTarget())
		return &c
	}()
	return p.run(ctx, o)
}

func (p *Process) run(ctx planProcessContext, o PlanOption) error {
	cmp, err := compose.GetAllCompose(ctx)
	if err != nil {
		return errs.ErrComposeNotFound
	}
	orcaplan := plan.BuildOrcaPlan(ctx, cmp)
	overlayer := compose.ComposeOverlayer(ctx, cmp)
	if ctx.Config().Volume.Enabled() {
		overlayer.OverlayVolume(orcaplan.Volumes)
	}
	if ctx.Config().Network.Enabled {
		overlayer.OverlayNetwork(orcaplan.Networks)
	}
	dotorca.DumpComposes(ctx, cmp, o.Force)
	pl, _ := yaml.Marshal(orcaplan)
	executor.NewFilesystem(p.ctx).WriteFile(p.ctx.OrcaPlanFile(), pl)
	return nil
}
