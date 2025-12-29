package compose

import (
	"orca/internal/compose/collector"
	"orca/internal/compose/dumper"
	inspector "orca/internal/inspector/compose"
	"orca/internal/policy"
)

func ComposeCollector(root string) collector.ComposeCollector {
	return collector.NewCollector(root, inspector.NewInspector(root))
}
func ComposeDumper(root string, p policy.ExecPolicy) dumper.ComposeDumper {
	return dumper.NewDumper(root, p)
}
