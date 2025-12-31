package docker

import (
	"fmt"
	"orca/errs"
	"orca/internal/context"
	"orca/model/policy"
	"os/exec"
)

type Executor interface {
	ComposeUp(string) (string, error)
	ComposeDown(string) (string, error)
	CreateNetwork(string, ...string) (string, error)
	CreateVolume(string, ...string) (string, error)
}

var _ Executor = (*executor)(nil)

func NewExecutor(p policy.ExecPolicy) Executor {
	return &executor{WithPolicy: context.NewWithPolicy(p)}
}

type executor struct {
	context.WithPolicy
}



func (d *executor) ComposeUp(composeFile string) (string, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	return d.run(cmd)
}

// docker compose down -f <file>
func (d *executor) ComposeDown(composeFile string) (string, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "down")
	return d.run(cmd)
}

// docker network create <name> [opt...]
func (d *executor) CreateNetwork(name string, opt ...string) (string, error) {
	c := append([]string{"network", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}

// docker volume create <name> [opt...]
func (d *executor) CreateVolume(name string, opt ...string) (string, error) {
	c := append([]string{"volume", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}

func (d *executor) run(cmd *exec.Cmd) (string, error) {
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