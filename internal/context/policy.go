package context

import "orca/internal/policy"

type WithPolicy interface {
	Policy() policy.ExecPolicy
}
type withPolicy struct {
	policy policy.ExecPolicy
}

func NewWithPolicy(p policy.ExecPolicy) withPolicy {
	if p == nil {
		panic("ExecPolicy must not be nil")
	}
	return withPolicy{policy: p}
}

func (w withPolicy) Policy() policy.ExecPolicy {
	return w.policy
}

type F struct {
	WithPolicy
}
