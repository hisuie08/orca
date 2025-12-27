package applier

import (
	"fmt"
	"orca/consts"
	"orca/ostools"
	"path/filepath"
)

type ComposeWriter interface {
	DumpCompose(name string, content []byte) (string, error)
}

type DotOrcaDumper struct {
	OrcaRoot string
}

func (d DotOrcaDumper) DumpCompose(
	name string, content []byte) (string, error) {
	cmpName := fmt.Sprintf("compose.%s.yml", name)
	target := filepath.Join(d.OrcaRoot, consts.DotOrcaDir, cmpName)
	return ostools.CreateFile(target, content, false)
}
