package context

import (
	"io"
	"os"

	"golang.org/x/term"
)

var _ WithColor = (*withColor)(nil)

type WithColor interface {
	Enabled() bool
}

type withColor struct {
	enabled bool
}


func (c *withColor) Enabled() bool {
	return c.enabled
}

func isTTY(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}
