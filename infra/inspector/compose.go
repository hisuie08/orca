package inspector

import (
	"orca/errs"
	"os"
	"os/exec"
	"path/filepath"
)

var _ ComposeInspector = (*dockerComposeInspector)(nil)

type ComposeInspector interface {
	Root() string
	Config(string) ([]byte, error)
}

func Compose(orcaRoot string) *dockerComposeInspector {
	return &dockerComposeInspector{orcaRoot: orcaRoot}

}

// dockerComposeInspector
type dockerComposeInspector struct {
	orcaRoot string
}

func (d dockerComposeInspector) Root() string {
	return d.orcaRoot
}
func (d dockerComposeInspector) Config(composeDir string) ([]byte, error) {
	// HACK: 駆け上がり探索防止用の空compose
	stopper := filepath.Join(d.orcaRoot, "compose.yml")
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
