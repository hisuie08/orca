package orca

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)
// 色文字用
type StringColor string

const (
	ColorReset  StringColor = "\033[0m"
	ColorRed    StringColor = "\033[31m"
	ColorGreen  StringColor = "\033[32m"
	ColorYellow StringColor = "\033[33m"
	ColorBlue   StringColor = "\033[34m"
	ColorGray   StringColor = "\033[90m"
)

type Colorizer struct {
	Enabled bool
}

func NewColorizer(w io.Writer) Colorizer {
	return Colorizer{Enabled: IsTTY(w)}
}
func (c Colorizer) Color(s string, col StringColor) string {
	if !c.Enabled {
		return s
	}
	return fmt.Sprintf("%s%s%s", col, s, ColorReset)
}

func (c Colorizer) Red(s string) string {
	return c.Color(s, ColorRed)
}

func (c Colorizer) Yellow(s string) string {
	return c.Color(s, ColorYellow)
}

func (c Colorizer) Green(s string) string {
	return c.Color(s, ColorGreen)
}

func (c Colorizer) Gray(s string) string {
	return c.Color(s, ColorGray)
}

func IsTTY(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}
