package policy

type ExecPolicy interface {
	AllowSideEffect() bool
}

var _ ExecPolicy = (*execPolicy)(nil)

type execPolicy struct {
	allowSideEfect bool
}

func (e *execPolicy) AllowSideEffect() bool {
	return e.allowSideEfect
}

var (
	Real = &execPolicy{allowSideEfect: true}
	Dry  = &execPolicy{allowSideEfect: false}
)
