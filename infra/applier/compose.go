package applier

import (
	"fmt"
	"orca/ostools"
	"path/filepath"
)

type ComposeFileWriter struct {
	path string
}

func (c ComposeFileWriter) WriteCompose(
	name string, content []byte) (string, error) {
	n := fmt.Sprintf("compose.%s.yml", name)
	p := filepath.Join(c.path, n) // orcaRoot/.orca/compose.
	return ostools.CreateFile(p, content)
}

type dryComposeWriter struct {
	path string
}

func (d dryComposeWriter) WriteCompose(
	name string, b []byte) (string, error) {
	n := fmt.Sprintf("compose.%s.yml", name)
	p := filepath.Join(d.path, n)
	return p, nil
}
