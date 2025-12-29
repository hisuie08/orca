package compose

import (
	"orca/errs"
	"orca/internal/context"
	"orca/internal/inspector/fs"
	"orca/model/compose"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var _ ComposeInspector = (*composeInspector)(nil)

type ComposeInspector interface {
	Config(string) (*compose.ComposeSpec, error)
}

func NewInspector(root string) ComposeInspector {
	return &composeInspector{
		WithRoot: context.NewWithRoot(root),
		fs:       fs.NewFsInspector,
	}

}

// composeInspector
type composeInspector struct {
	context.WithRoot
	fs fs.FsInspector
}

func (d *composeInspector) Config(composeDir string) (*compose.ComposeSpec, error) {
	// HACK: 駆け上がり探索防止用の空compose
	stopper := filepath.Join(d.Root(), "compose.yml")
	created := false

	if _, err := os.Stat(stopper); os.IsNotExist(err) {
		os.WriteFile(stopper, []byte{}, 0644)
		created = true
	}

	cmd := exec.Command(
		"docker", "compose",
		"--project-directory", composeDir,
		"config",
	)
	out, err := cmd.CombinedOutput()

	if created {
		os.Remove(stopper)
	}

	if err != nil {
		return nil, errs.ErrComposeNotFound
	}
	spec := &compose.ComposeSpec{}
	if err := yaml.Unmarshal(out, spec); err != nil {
		return nil, err
	}
	return spec, nil
}
