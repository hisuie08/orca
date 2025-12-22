package applier

import (
	"orca/consts"
	"orca/internal/context"
	"orca/ostools"
	"path/filepath"
)

type ComposeWriter interface {
	DumpCompose(name string, content []byte) (string, error)
}

type DotOrcaDumper struct {
	Ctx context.OrcaContext
}

func (d DotOrcaDumper) DumpCompose(
	name string, content []byte) (string, error) {
		cmpName:="compose."+name+".yml"
	target := filepath.Join(d.Ctx.OrcaRoot, consts.DotOrcaDir, cmpName)
	switch d.Ctx.RunMode {
	case context.ModeExecute:
		if err := ostools.CreateFile(target, content); err != nil {
			return "", err
		}
	case context.ModeDryRun:
		return target, nil

	}
	return target, nil
}
