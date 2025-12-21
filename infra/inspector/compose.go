package inspector

import (
	"errors"
	"orca/internal/ostools"
	"os"
	"os/exec"
	"path/filepath"
)

var ErrNoCompose = errors.New("no compose")

// DockerComposeInspector
type DockerComposeInspector struct {
	OrcaRoot string
}

func (d DockerComposeInspector) Directories() ([]string, error) {
	return ostools.Directories(d.OrcaRoot)
}

func (d DockerComposeInspector) Config(composeDir string) ([]byte, error) {
	// HACK: 駆け上がり探索防止用の空compose
	stopper := filepath.Join(d.OrcaRoot, "compose.yml")
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
		return nil, ErrNoCompose
	}
	return out, nil
}
