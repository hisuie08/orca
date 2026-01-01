package context

import (
	"path/filepath"
)

type WithRoot interface {
	Root() string
	OrcaDir() string
	OrcaYamlFile() string
}
type withRoot struct {
	root string
}

func (w *withRoot) Root() string {
	return w.root
}

func (w *withRoot) OrcaDir() string {
	return filepath.Join(w.root, ".orca")
}

func (w *withRoot) OrcaYamlFile() string {
	return filepath.Join(w.root, "orca.yml")
}
