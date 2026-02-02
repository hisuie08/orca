package internal

import (
	"fmt"
	"orca/model/config"
	"orca/model/plan"
)

const (
	cluster = "orca.cluster"
	compose = "orca.compose"

	volType   = "orca.volume.type"
	volLocal  = "local"
	volShared = "shared"
)

func commonLabel(cfg config.OrcaConfig) []string {
	return []string{label(cluster, cfg.Name)}
}
func label(k string, v any) string {
	return fmt.Sprintf("%s=%v", k, v)
}

func VolumeLabel(cfg config.OrcaConfig, vp plan.VolumePlan) []string {
	labels := commonLabel(cfg)
	labels = append(labels, label(volType, vp.Type))
	return labels
}
