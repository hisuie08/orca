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
var _ WithOutput = (*Context)(nil)
var _ WithReport = (*Context)(nil)

type Context struct {
	root   *withRoot
	config *withConfig
	policy *withPolicy
	color  *withColor
	output *withOutput
	report *withReport
}

func New() Context {
	return Context{}
}

func FromCommandCtx(ctx CommandContext) Context {
	c := New().WithRoot(ctx.Root()).WithOutput(ctx.Output()).
		WithReport(ctx.Report())
	return c
}

func (c Context) WithRoot(root string) Context {
	c.root = newWithRoot(root)
	return c
}
func (c Context) WithConfig(cfg *config.ResolvedConfig) Context {
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

func (c Context) WithOutput(w io.Writer) Context {
	c.output = &withOutput{out: w}
	return c
}

func (c Context) WithReport(w io.Writer) Context {
	c.report = &withReport{out: w}
	return c
}

func (c Context) WithColor(w io.Writer) Context {
	c.color = &withColor{enabled: isTTY(w)}
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

func (c *Context) Config() *config.ResolvedConfig {
	return c.config.Config()
}

func (c *Context) Policy() policy.ExecPolicy {
	return c.policy.Policy()
}

func (c *Context) Colored() bool {
	return c.color.Colored()
}

func (c *Context) Output() io.Writer {
	return c.output.out
}
func (c *Context) Report() io.Writer {
	return c.report.out
}
