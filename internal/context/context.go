package context

import (
	"fmt"
	"io"
	"orca/model/config"
	"orca/model/policy"
	"path/filepath"
)

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
	if root == "" {
		panic("orca root must not be empty")
	}

	abs, err := filepath.Abs(root)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve absolute path: %v", err))
	}
	c.root = &withRoot{root: abs}
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

func (c Context) WithColor(w io.Writer) Context {
	c.color = &withColor{enabled: isTTY(w)}
	return c
}
