package compose

import (
	"orca/errs"
	"orca/internal/context"
	"os"
	"os/exec"
	"path/filepath"
)

var _ ComposeInspector = (*dockerComposeInspector)(nil)

type ComposeInspector interface {
	Config(string) ([]byte, error)
}

func NewInspector(orcaRoot string) ComposeInspector {
	return &dockerComposeInspector{WithRoot: context.NewWithRoot(orcaRoot)}

}

// dockerComposeInspector
type dockerComposeInspector struct {
	context.WithRoot
}

func (d *dockerComposeInspector) Config(composeDir string) ([]byte, error) {
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
	return out, nil
}
