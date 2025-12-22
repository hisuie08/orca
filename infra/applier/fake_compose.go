package applier

import (
	"orca/consts"
	"path/filepath"
)

type FakeDotOrcaDumper struct {
	FakeRoot string
	FakeDir  map[string][]byte
}

func (d FakeDotOrcaDumper) DumpCompose(
	name string, content []byte) (string, error) {
	cmpName := "compose." + name + ".yml"
	d.FakeDir[cmpName] = content
	return filepath.Abs(filepath.Join(d.FakeRoot, consts.DotOrcaDir, name))
}
