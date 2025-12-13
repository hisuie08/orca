package ostools

import (
	"fmt"
	orca "orca/helper"
	"os/exec"
)

// コマンド実行用
func run(c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	out, err := cmd.CombinedOutput()
	op := fmt.Sprintf("%s %v failed: ", c, args)
	return string(out), orca.OrcaError(op, err)
}

// Volume
//
// docker volume inspect <name>
func VolumeExists(name string) bool {
	_, err := run("docker", "volume", "inspect", name)
	return err == nil
}

// Network
//
// docker network inspect <name>
func NetworkExists(name string) bool {
	_, err := run("docker", "network", "inspect", name)
	return err == nil
}

// docker network create <name> <--internal>
func CreateNetwork(name string, internal bool) error {
	cmd := []string{"network", "create", name}
	if internal {
		cmd = append(cmd, "--internal")
	}
	_, err := run("docker", cmd...)
	return err
}

// Compose
//
// docker compose --project-directory <dir> config
func ComposeConfig(dir string) (string, error) {
	return run("docker", "compose", "--project-directory", dir, "config")
}

// docker compose up -d -f <file>
func ComposeUp(composeFile string) error {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compose up failed: %v: %s", err, out)
	}
	return nil
}

// docker compose down -f <file>
func ComposeDown(composeFile string, workdir string) error {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "down")
	cmd.Dir = workdir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compose down failed: %v: %s", err, out)
	}
	return nil
}
