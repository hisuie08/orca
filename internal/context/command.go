package context

type CommandContext interface {
	WithOutput
	WithReport
	WithRoot
}

func FromCommandCtx(ctx CommandContext) Context {
	c := New().WithRoot(ctx.Root()).WithOutput(ctx.Output()).
		WithReport(ctx.Report())
	return c
}
