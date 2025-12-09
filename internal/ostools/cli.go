package ostools

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Dockerコマンド実行用
func run(args ...string) (string, error) {
	cmd := exec.Command("docker", args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("docker %v failed: %v: %s",
			args, err, stderr.String())
	}

	return out.String(), nil
}

// Volume
// 
func VolumeExists(name string) bool {
	_, err := run("volume", "inspect", name)
	return err == nil
}

// Network
func NetworkExists(name string) bool {
	_, err := run("network", "inspect", name)
	return err == nil
}

func CreateNetwork(name string, driver string) error {
	if driver == "" {
		driver = "bridge"
	}

	_, err := run("network", "create", "--driver", driver, name)
	return err
}

// Compose

// compose up -d -f <file>
func ComposeUp(composeFile string, workdir string) error {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	cmd.Dir = workdir
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
