package volume

import (
	"fmt"
	"orca/internal/context"
	"orca/internal/executor"
	. "orca/internal/usecase/docker/internal"
	"orca/model/plan"
)

type createVolContext interface {
	context.WithPolicy
	context.WithConfig
	context.WithLog
}
type dockerExecutor interface {
	CreateVolume(name string, opts ...string) ([]byte, error)
}

func CreateVolume(ctx createVolContext, vp plan.VolumePlan) {
	de := executor.NewDocker(ctx)
	createVolume(ctx, vp, de)
}

func buildOption(vp plan.VolumePlan) []string {
	return []string{"-o type=none", "-o o=bind",
		fmt.Sprintf("-o device=%s", vp.BindPath)}
}
func createVolume(ctx createVolContext, vp plan.VolumePlan, de dockerExecutor) {
	opts := append(buildOption(vp), VolumeLabel(*ctx.Config(), vp)...)
	de.CreateVolume(vp.Name, opts...)
}
