package compose

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

var ErrNoCompose = errors.New("no compose")

type ComposeInspector interface {
	Config(dir string) ([]byte, error)
}
type DockerInspector struct{}

type DockerComposeInspector struct {
	Root string
}

func (d DockerComposeInspector) Config(dir string) ([]byte, error) {
	stopper := filepath.Join(d.Root, "compose.yml")
	created := false

	if _, err := os.Stat(stopper); os.IsNotExist(err) {
		os.WriteFile(stopper, []byte{}, 0644)
		created = true
	}

	cmd := exec.Command(
		"docker", "compose",
		"--project-directory", dir,
		"config",
	)
	out, err := cmd.CombinedOutput()

	if created {
		os.Remove(stopper)
	}

	if err != nil {
		return nil, err
	}
	return out, nil
}

type FakeInspector struct {
	Results map[string][]byte
}

func (f FakeInspector) Config(dir string) ([]byte, error) {
	v, ok := f.Results[dir]
	if !ok {
		return nil, errors.New("no compose")
	}
	return v, nil
}
