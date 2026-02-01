package plansummary

import (
	"fmt"
	"orca/model/plan"
	"orca/presenter/formatter/printer"
	"strings"
)

func fmtVolPlan(vp plan.VolumePlan, o ViewOption) []string {
	name := vp.Name
	typ := func() string {
		typ := string(vp.Type)
		// switch vp.Type {
		// case plan.VolumeShared:
		// 	typ = color.ColorString("shared", color.Blue, o.Colored())
		// case plan.VolumeLocal:
		// 	typ = color.ColorString("local", color.Green, o.Colored())
		// case plan.VolumeExternal:
		// 	typ = color.ColorString("external", color.Gray, o.Colored())
		// }
		return typ
	}()
	ub := func() string {
		result := []string{}
		for _, vr := range vp.UsedBy {
			result = append(result, fmt.Sprintf("%s(%s)", vr.Compose, vr.Key))
		}
		return strings.Join(result, ",")
	}()
	bind := func() string {
		if vp.BindPath == "" {
			return "-"
		}
		return vp.BindPath
	}()
	status := fmt.Sprintf("%t / %t", vp.Exists, vp.NeedMkdir)
	return []string{name, typ, ub, bind, status}
}

func PrintVolumePlanTable(plans []plan.VolumePlan, o ViewOption) string {
	title := "[VOLUME PLAN]"
	headers := []string{"NAME", "TYPE", "USED BY", "BIND PATH", "EXISTS(VOL/PATH)"}
	rows := make([][]string, 0, len(plans))
	for _, p := range plans {
		row := fmtVolPlan(p, o)
		if len(row) != len(headers) {
			panic("PrintVolumePlanTable got invalid data")
		}
		rows = append(rows, row)
	}
	return printer.PTable(title, headers, rows)
}
