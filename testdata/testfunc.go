package testdata
// 実験用コード片置き場
import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func SpewPrint(c ...any) {
	spew.Config.DisableCapacities = true
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisablePointerMethods = true
	spew.Dump(c...)
}

const (
	action.Message = "default network is overridden to use shared network orca_network"
	action.Message = fmt.Sprintf(
			"network %s conflicts with shared network and will be removed",
			net.Spec.Name)
)

const (
	// Warning 判定
		// パスが存在しないのにensure_path=falseで作成が許可されていないとき
		switch {
		case vp.Type != plan.VolumeExternal &&
			vp.NeedMkdir &&
			!cfg.EnsurePath:
			warningMsg := fmt.Sprintf("bind path %s does not exist and ensure_path=false", vp.BindPath)
			vp.Warnings = append(vp.Warnings, warningMsg)
		case vp.Type == plan.VolumeExternal && !di.VolumeExists(vp.Name):
			// externalだけどplan時点でボリュームが存在しないとき
			warningMsg := fmt.Sprintf("external volume %s does not exist", vp.Name)
			vp.Warnings = append(vp.Warnings, warningMsg)
		}
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

func PrintNetworkPlan(p NetworkPlan, printer *orca.Printer) {
	title := "[NETWORK PLAN]"
	printer.Printf("%s\n", title)
	printer.Printf("SHARED NETWORK: %s\n", p.SharedName)
	printer.Printf("Compose Changes: %d\n", len(p.Actions))

	// compose名でソート
	composes := make([]string, 0, len(p.Actions))
	for k := range p.Actions {
		composes = append(composes, k)
	}
	sort.Strings(composes)

	for _, compose := range composes {
		actions := p.Actions[compose]
		if len(actions) == 0 {
			continue
		}

		printer.Printf("%s\n", compose)

		for _, a := range actions {
			switch a.Type {
			case NetworkOverrideDefault:
				label := printer.C.Blue("override")
				printer.Printf("  - %s %s → %s\n", label, a.Network, p.SharedName)
			case NetworkRemoveConflict:
				label := printer.C.Yellow("remove")
				printer.Printf("  - %s %s (name conflict)\n", label, a.Network)
			}
		}
	}
	printer.Printf("")
}


func indent(n int, s string) string {
	pad := strings.Repeat(" ", n)
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		if l != "" {
			lines[i] = pad + l
		}
	}
	return strings.Join(lines, "\n")
}


func sortByName[V any](s []V, k func(V) string) {
	slices.SortStableFunc(s, func(i, j V) int {
		return strings.Compare(k(i), k(j))
	})
}
