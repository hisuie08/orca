package context

import (
	"io"
	"orca/model/config"
	"orca/model/policy"
	"orca/model/policy/log"
)

var _ WithRoot = (*Context)(nil)
var _ WithConfig = (*Context)(nil)
var _ WithPolicy = (*Context)(nil)
var _ WithColor = (*Context)(nil)
var _ WithLog = (*Context)(nil)

type Context struct {
	root   *withRoot
	config *withConfig
	policy *withPolicy
	color  *withColor
	log    *withLog
}

func New() Context {
	return Context{}
}

func (c Context) WithRoot(root string) Context {
	c.root = newWithRoot(root)
	return c
}
func (c Context) WithConfig(cfg *config.OrcaConfig) Context {
	if cfg == nil {
		panic("ResolvedConfig must not be nil")
	}
	c.config = &withConfig{config: cfg}
	return c
}
func (c Context) WithPolicy(p policy.ExecPolicy) Context {
	if p == nil {
		panic("ExecPolicy must not be nil")
	}
	c.policy = &withPolicy{policy: p}
	return c
}

func (c Context) WithColor(w io.Writer) Context {
	c.color = &withColor{enabled: isTTY(w)}
	return c
}

func (c Context) WithLog(l log.LogLevel, o io.Writer) Context {
	c.log = &withLog{logLevel: l, out: o}
	return c
}

func (c *Context) Root() string {
	return c.root.Root()
}

func (c *Context) OrcaDir() string {
	return c.root.OrcaDir()
}
func (c *Context) OrcaYamlFile() string {
	return c.root.OrcaYamlFile()
}

func (c *Context) Config() *config.OrcaConfig {
	return c.config.Config()
}

func (c *Context) Policy() policy.ExecPolicy {
	return c.policy.Policy()
}

func (c *Context) Colored() bool {
	return c.color.Colored()
}

func (c *Context) LogLevel() log.LogLevel {
	return c.log.LogLevel()
}

func (c *Context) LogTarget() io.Writer {
	return c.log.out
}
