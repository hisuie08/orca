package context

import (
	"fmt"
	"path/filepath"
)

type WithRoot struct {
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

	return WithRoot{root: abs}
}

func (w WithRoot) Root() string {
	return w.root
}

func (w WithRoot) OrcaDir() string {
	return filepath.Join(w.root, ".orca")
}

func (w WithRoot) OrcaYamlFile() string {
	return filepath.Join(w.root, "orca.yml")
}
