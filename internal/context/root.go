package context

import (
	"fmt"
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

func NewWithRoot(root string) WithRoot {
	if root == "" {
		panic("orca root must not be empty")
	}

	abs, err := filepath.Abs(root)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve absolute path: %v", err))
	}

	return &withRoot{root: abs}
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
