package internal

import (
	"fmt"
	"orca/model/config"
	"orca/model/plan"
	"strings"
)

const (
	managed = "orca.managed"
	cluster = "orca.cluster"
	compose = "orca.compose"

	volType   = "orca.volume.type"
	volLocal  = "local"
	volShared = "shared"
)

func commonLabel(cfg config.OrcaConfig) []string {
	return []string{label(managed, true),
		label(cluster, cfg.Name)}
}
func label(k string, v any) string {
	return fmt.Sprintf("%s=%v", k, v)
}

func VolumeLabel(cfg config.OrcaConfig, vp plan.VolumePlan) []string {
	labels := commonLabel(cfg)
	labels = append(labels, label(volType, vp.Type),
		label(compose, usedBy(vp.UsedBy)))
	return labels
}

func usedBy(vrs []plan.VolumeRef) string {
	refs := []string{}
	for _, vr := range vrs {
		refs = append(refs, fmt.Sprintf("%s", vr.Compose))
	}
	return fmt.Sprintf("[%s]", strings.Join(refs, ", "))
}
