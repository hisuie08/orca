package ostools

import (
	"fmt"
	"os/exec"
)

// Compose
//

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
