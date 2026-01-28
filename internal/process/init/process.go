package pinit

import (
	"orca/internal/context"
	"orca/internal/logger"
	"orca/internal/usecase/config"
	. "orca/model/config"
	"orca/model/policy/log"
)

type InitOption struct {
	CfgOption
	config.WriteOption
}

type initProcessContext interface {
	context.CommandContext
}

type InitProcess struct {
	ctx    context.CommandContext
	logger logger.Logger
}

func New(c context.CommandContext) *InitProcess {
	return &InitProcess{ctx: c, logger: logger.New(c)}
}

func (p *InitProcess) Run(opt InitOption) error {
	ctx := p.ctx
	return p.run(ctx, opt)
}

func (p *InitProcess) run(ctx initProcessContext, opt InitOption) error {
	p.logger.Logln(log.LogNormal, "initializing orca cluster")
	cfg := config.Create(ctx, opt.CfgOption)
	if err := config.Write(ctx, cfg, opt.WriteOption); err != nil {
		return err
	}
	return nil
}
