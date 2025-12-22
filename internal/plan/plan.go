package plan

import (
	"io"
	"orca/internal/context"
)

func DumpPlan(
	ctx context.OrcaContext,
	volumePlans []VolumePlan,
	networkPlan NetworkPlan,
	w io.Writer,
) {
	ctx.Printer.SetWriter(w)
	ctx.Printer.Printf("[Orca Config]\n")
	ctx.Printer.Printf("\n")
	PrintVolumePlanTable(volumePlans, ctx.Printer)
	ctx.Printer.Printf("\n")
	PrintNetworkPlan(networkPlan, ctx.Printer)
}
