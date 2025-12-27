package applier

import (
	"orca/consts"
	"path/filepath"
)

var (
	Applier    *applier    = &applier{}
	DryApplier *dryApplier = &dryApplier{}
)

type applier struct{}

func (a applier) Compose(orcaRoot string) *ComposeFileWriter {
	p := filepath.Join(orcaRoot, consts.DotOrcaDir) // orcaRoot/.orca/
	return &ComposeFileWriter{path: p}
}

func (a applier) ConfigWriter(orcaRoot string) *configFileWriter {
	return &configFileWriter{OrcaRoot: orcaRoot}
}

type dryApplier struct{}

func (d dryApplier) Compose(orcaRoot string) *dryComposeWriter {
	p := filepath.Join(orcaRoot, consts.DotOrcaDir) // orcaRoot/.orca/
	return &dryComposeWriter{path: p}
}

func (d dryApplier) ConfigWriter(orcaRoot string) *dryConfigWriter {
	return &dryConfigWriter{OrcaRoot: orcaRoot}
}
