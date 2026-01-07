package context

import "io"

type ViewContext interface {
	WithColor
	WithOutput
	WithReport
	WithDiag
}
type WithOutput interface {
	Output() io.Writer
}

type WithReport interface {
	Report() io.Writer
}

type WithDiag interface {
	Diag() io.Writer
}

type withOutput struct {
	out io.Writer
}

func (w *withOutput) Output() io.Writer {
	return w.out
}

type withReport struct {
	out io.Writer
}

func (w *withReport) Report() io.Writer {
	return w.out
}

type withDiag struct {
	out io.Writer
}

func (w withDiag) Diag() io.Writer {
	return w.out
}
