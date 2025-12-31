package docker

import (
	"orca/errs"
	"orca/internal/inspector/filesystem"
	"orca/model/compose"
	"os/exec"

	"gopkg.in/yaml.v3"
)

var _ Inspector = (*inspector)(nil)

type Inspector interface {
	NetworkExists(name string) bool
	VolumeExists(name string) bool
	BindExists(dir string) bool
	Compose(dir string) (*compose.ComposeSpec, error)
}

// DockerInspector 実装
type inspector struct {
	fi filesystem.Inspector
}

func NewInspector() Inspector {
	return &inspector{
		fi: filesystem.NewInspector(),
	}
}

// VolumeExists run docker volume inspect <name>
func (d *inspector) VolumeExists(name string) bool {
	cmd := exec.Command("docker", "volume", "inspect", name)
	return d.inspectExists(cmd)
}

// NetworkExists run docker network inspect <name>
func (d *inspector) NetworkExists(name string) bool {
	cmd := exec.Command("docker", "network", "inspect", name)
	return d.inspectExists(cmd)
}

// BindExists returns boolean indicating whether the directory exists
// Delegates to filesystem.Inspector
func (d *inspector) BindExists(path string) bool {
	return d.fi.DirExists(path)
}

// Compose returns ErrComposeNotFound when docker compose config cannot be generated.
// This includes missing or invalid compose files.
func (d *inspector) Compose(dir string) (*compose.ComposeSpec, error) {
	cmd := exec.Command("docker", "compose", "--project-directory", dir, "config")
	o, err := d.runDocker(cmd)
	if err != nil {
		// config が生成できない = compose が無い or 無効
		return nil, errs.ErrComposeNotFound
	}
	spec := &compose.ComposeSpec{}
	if err := yaml.Unmarshal(o, spec); err != nil {
		return nil, err
	}
	return spec, nil
}

func (d *inspector) runDocker(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

func (d *inspector) inspectExists(cmd *exec.Cmd) bool {
	_, err := d.runDocker(cmd)
	if err == nil {
		return true
	}
	// docker inspect が "not found" を返すケースだけ false
	// それ以外は「存在確認できない」＝ false とする（がログ候補）
	return false
}
