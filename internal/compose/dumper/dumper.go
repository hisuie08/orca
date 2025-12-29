package dumper

import (
	"fmt"
	"orca/internal/context"
	"orca/internal/policy"
	"orca/ostools"
	"path/filepath"
)

var _ ComposeDumper = (*composeDumper)(nil)

type ComposeDumper interface {
	Write(string, []byte) (string, error)
}
type composeDumper struct {
	context.WithPolicy
	context.WithRoot
}

func NewDumper(root string, p policy.ExecPolicy) *composeDumper {
	return &composeDumper{
		WithRoot:   context.NewWithRoot(root),
		WithPolicy: context.NewWithPolicy(p),
	}
}

func (c *composeDumper) Write(
	name string, content []byte) (string, error) {
	n := fmt.Sprintf("compose.%s.yml", name)
	p := filepath.Join(c.OrcaDir(), n)
	if c.Policy().AllowSideEffect() {
		return ostools.CreateFile(p, content)
	}
	return p, nil
}
