package policy

type ExecPolicy interface {
	AllowSideEffect() bool
}

var _ ExecPolicy = (*RealPolicy)(nil)
var _ ExecPolicy = (*DryPolicy)(nil)

type RealPolicy struct{}

func (RealPolicy) AllowSideEffect() bool { return true }

type DryPolicy struct{}

func (DryPolicy) AllowSideEffect() bool { return false }
