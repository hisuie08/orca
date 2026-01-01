package context

import "orca/model/policy"

type WithPolicy interface {
	Policy() policy.ExecPolicy
}
type withPolicy struct {
	policy policy.ExecPolicy
}

func (w withPolicy) Policy() policy.ExecPolicy {
	return w.policy
}
