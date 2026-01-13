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
	context.WithReport
}

func DumpComposes(
	ctx DumpContext, cm compose.ComposeMap, force bool) ([]string, error) {
	return dump.DotOrcaDumper(ctx, force).DumpComposes(cm)
}

func DumpPlan(ctx DumpContext, pl plan.OrcaPlan, force bool) (string, error) {
	return dump.DotOrcaDumper(ctx, force).DumpPlan(pl)
}
