package dotorca

import (
	"orca/internal/context"
	"orca/internal/usecase/dotorca/dump"
	"orca/model/compose"
	"orca/model/plan"
)

// .orca/
type DumpContext interface {
	context.WithRoot
	context.WithPolicy
}

type Dumper interface {
	DumpComposes(compose.ComposeMap) ([]string, error)
	DumpPlan(plan.OrcaPlan) (string, error)
}

func DotOrcaDumper(ctx DumpContext, force bool) Dumper {
	return dump.DotOrcaDumper(ctx, force)
}
