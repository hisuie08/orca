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
	// shared volumeだけorcaが自前で作る
	// それ以外はdockerがcomposeから作るか、external
	if !vp.Exists && vp.Type == plan.VolumeShared {
		createVolume(ctx, vp, de)
	}

}

func createVolume(ctx createVolContext, vp plan.VolumePlan, de dockerExecutor) {
	opts := []string{}
	if vp.BindPath != "" {
		opts = append(opts, "-o type=none", "-o o=bind",
			fmt.Sprintf("-o device=%s", vp.BindPath))
	}
	opts = append(opts, VolumeLabel(*ctx.Config(), vp)...)
	de.CreateVolume(vp.Name, opts...)
}
