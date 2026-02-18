package capability

import (
	"io"
	"orca/model/config"
	"orca/model/policy"
	"orca/model/policy/log"
)

var _ WithRoot = (*Capability)(nil)
var _ WithConfig = (*Capability)(nil)
var _ WithPolicy = (*Capability)(nil)
var _ WithColor = (*Capability)(nil)
var _ WithLog = (*Capability)(nil)

type Capability struct {
	root   *withRoot
	config *withConfig
	policy *withPolicy
	color  *withColor
	log    *withLog
}

func New() Capability {
	return Capability{}
}

func (c Capability) WithRoot(root string) Capability {
	c.root = newWithRoot(root)
	return c
}
func (c Capability) WithConfig(cfg *config.OrcaConfig) Capability {
	if cfg == nil {
		panic("ResolvedConfig must not be nil")
	}
	c.config = &withConfig{config: cfg}
	return c
}
func (c Capability) WithPolicy(p policy.ExecPolicy) Capability {
	if p == nil {
		panic("ExecPolicy must not be nil")
	}
	c.policy = &withPolicy{policy: p}
	return c
}

func (c Capability) WithColor(w io.Writer) Capability {
	c.color = &withColor{enabled: isTTY(w)}
	return c
}

func (c Capability) WithLog(l log.LogLevel, o io.Writer) Capability {
	c.log = &withLog{logLevel: l, out: o}
	return c
}

func (c *Capability) Root() string {
	return c.root.Root()
}

func (c *Capability) OrcaDir() string {
	return c.root.OrcaDir()
}
func (c *Capability) OrcaYamlFile() string {
	return c.root.OrcaYamlFile()
}
func (c *Capability) OrcaPlanFile() string {
	return c.root.OrcaPlanFile()
}
func (c *Capability) Config() *config.OrcaConfig {
	return c.config.Config()
}

func (c *Capability) Policy() policy.ExecPolicy {
	return c.policy.Policy()
}

func (c *Capability) Colored() bool {
	return c.color.Colored()
}

func (c *Capability) LogLevel() log.LogLevel {
	return c.log.LogLevel()
}

func (c *Capability) LogTarget() io.Writer {
	return c.log.out
}
