package context

import (
	"io"
	"orca/model/config"
	"orca/model/policy"
)

var _ WithRoot = (*Context)(nil)
var _ WithConfig = (*Context)(nil)
var _ WithPolicy = (*Context)(nil)
var _ WithColor = (*Context)(nil)

type Context struct {
	root   *withRoot
	config *withConfig
	policy *withPolicy
	color  *withColor
}

func New() Context {
	return Context{}
}

func (c Context) WithRoot(root string) Context {
	c.root = newWithRoot(root)
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

func (c Context) WithConfig(cfg *config.ResolvedConfig) Context {
	if cfg == nil {
		panic("ResolvedConfig must not be nil")
	}
	c.config = &withConfig{config: cfg}
	return c
}

func (c *Context) Config() *config.ResolvedConfig {
	return c.config.Config()
}

func (c Context) WithPolicy(p policy.ExecPolicy) Context {
	if p == nil {
		panic("ExecPolicy must not be nil")
	}
	c.policy = &withPolicy{policy: p}
	return c
}

func (c *Context) Policy() policy.ExecPolicy {
	return c.policy.Policy()
}

func (c Context) WithColor(w io.Writer) Context {
	c.color = &withColor{enabled: isTTY(w)}
	return c
}

func (c *Context) Enabled() bool {
	return c.color.Enabled()
}
