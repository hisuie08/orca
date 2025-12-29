package docker

import (
	"fmt"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/policy"
	"os/exec"
)

type Executor interface {
	ComposeUp(string) (string, error)
	ComposeDown(string) (string, error)
	CreateNetwork(string, ...string) (string, error)
	CreateVolume(string, ...string) (string, error)
}

var _ Executor = (*dockerExecutor)(nil)

func NewExecutor(p policy.ExecPolicy) Executor {
	return &dockerExecutor{WithPolicy: context.NewWithPolicy(p)}
}

type dockerExecutor struct {
	context.WithPolicy
}

func (d *dockerExecutor) run(cmd *exec.Cmd) (string, error) {
	if d.Policy().AllowSideEffect() {
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(out), &errs.ExternalError{Cmd: cmd.String(), Err: err}
		} else {
			return string(out), nil
		}
	}
	out := fmt.Sprintf("[DRY-RUN] %s", cmd.String())
	return out, nil
}

func (d *dockerExecutor) ComposeUp(composeFile string) (string, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	return d.run(cmd)
}

// docker compose down -f <file>
func (d *dockerExecutor) ComposeDown(composeFile string) (string, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "down")
	return d.run(cmd)
}

// docker network create <name> [opt...]
func (d *dockerExecutor) CreateNetwork(name string, opt ...string) (string, error) {
	c := append([]string{"network", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}

// docker volume create <name> [opt...]
func (d *dockerExecutor) CreateVolume(name string, opt ...string) (string, error) {
	c := append([]string{"volume", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}
