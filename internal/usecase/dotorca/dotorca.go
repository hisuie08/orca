package dotorca

import (
	"orca/internal/capability"
	"orca/internal/usecase/dotorca/dump"
	"orca/model/compose"
	"orca/model/plan"
)

// .orca/
type DumpCapability interface {
	capability.WithRoot
	capability.WithPolicy
	capability.WithLog
}

func DumpComposes(
	caps DumpCapability, cm compose.ComposeMap, force bool) ([]string, error) {
	return dump.DotOrcaDumper(caps, force).DumpComposes(cm)
}

func DumpPlan(caps DumpCapability, pl plan.OrcaPlan, force bool) (string, error) {
	return dump.DotOrcaDumper(caps, force).DumpPlan(pl)
}
