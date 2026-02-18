package capability

import (
	"io"
	"os"

	"golang.org/x/term"
)

var _ WithColor = (*withColor)(nil)

type WithColor interface {
	Colored() bool
}

type withColor struct {
	enabled bool
}

func (c *withColor) Colored() bool {
	return c.enabled
}

func isTTY(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}
