package policy

type ExecPolicy interface {
	AllowSideEffect() bool
}

var (
	Real = &realPolicy{}
	Dry  = &dryPolicy{}
)

var _ ExecPolicy = (*realPolicy)(nil)
var _ ExecPolicy = (*dryPolicy)(nil)

type realPolicy struct{}

func (realPolicy) AllowSideEffect() bool { return true }

type dryPolicy struct{}

func (dryPolicy) AllowSideEffect() bool { return false }
