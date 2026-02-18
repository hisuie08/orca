package volume

import (
	"fmt"
	"orca/internal/capability"
	"orca/internal/executor"
	. "orca/internal/usecase/docker/internal"
	"orca/model/plan"
)

type createVolCapability interface {
	capability.WithPolicy
	capability.WithConfig
	capability.WithLog
}
type dockerExecutor interface {
	CreateVolume(name string, opts ...string) ([]byte, error)
}

func CreateVolume(caps createVolCapability, vp plan.VolumePlan) {
	de := executor.NewDocker(caps)
	// shared volumeだけorcaが自前で作る
	// それ以外はdockerがcomposeから作るか、external
	if !vp.Exists && vp.Type == plan.VolumeShared {
		createVolume(caps, vp, de)
	}

}

func createVolume(caps createVolCapability, vp plan.VolumePlan, de dockerExecutor) {
	opts := []string{}
	if vp.BindPath != "" {
		opts = append(opts, "-o type=none", "-o o=bind",
			fmt.Sprintf("-o device=%s", vp.BindPath))
	}
	opts = append(opts, VolumeLabel(*caps.Config(), vp)...)
	de.CreateVolume(vp.Name, opts...)
}
