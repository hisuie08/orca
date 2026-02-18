package pinit

import (
	"fmt"
	"orca/internal/capability"
	"orca/internal/logger"
	"orca/internal/usecase/config"
	. "orca/model/config"
	"orca/model/policy/log"
)

type InitOption struct {
	CfgOption
	config.WriteOption
}

type initProcessCapability interface {
	capability.CommandCapability
}

type InitProcess struct {
	caps   capability.CommandCapability
	logger logger.Logger
}

func New(c capability.CommandCapability) *InitProcess {
	return &InitProcess{caps: c, logger: logger.New(c)}
}

func (p *InitProcess) Run(opt InitOption) error {
	caps := p.caps
	return p.run(caps, opt)
}

func (p *InitProcess) run(caps initProcessCapability, opt InitOption) error {
	p.logger.Logln(log.LogNormal, "initializing orca cluster")
	cfg := config.Create(caps, opt.CfgOption)
	if err := config.Write(caps, cfg, opt.WriteOption); err != nil {
		return err
	}
	p.logger.Logln(log.LogNormal, fmt.Sprintf("cluster %s was initialized", opt.Name))
	return nil
}
