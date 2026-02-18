package network

import (
	"orca/internal/capability"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
	"testing"
)

var _ dockerInspector = (*fakeDockerInspector)(nil)

type fakeDockerInspector struct {
	Exists string
}

func (f *fakeDockerInspector) NetworkExists(name string) bool {
	return name == f.Exists
}
func fakeNetCaps(name string, enabled bool) NetworkPlanCapability {
	n := config.NetworkConfig{Name: name, Enabled: enabled, Internal: false}
	caps := capability.New().WithConfig(&config.OrcaConfig{Network: n})
	return &caps
}
func Test_CreateOrNot(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		exists  string
		want    bool
	}{
		{name: "create", enabled: true, exists: "", want: true},
		{name: "exists", enabled: true, exists: "exists", want: false},
		{name: "disabled", enabled: false, exists: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildNetworkPlan(
				fakeNetCaps(tt.name, tt.enabled),
				[]compose.CollectedNetwork{},
				&fakeDockerInspector{Exists: tt.exists})
			// Create = enabled && !exists
			if got.Create != tt.want {
				t.Errorf("incorrect decisions Create should be %t but got %t",
					tt.want, got.Create)
			}
		})
	}
}

func fakeNetwork(c, k, n string) compose.CollectedNetwork {
	return compose.CollectedNetwork{
		Ref:  compose.FromRef{Compose: c, Key: k},
		Spec: &compose.NetworkSpec{Name: n}}
}

func TestActionType(t *testing.T) {
	netname := "shared"
	tests := []struct {
		name    string
		network compose.CollectedNetwork
		want    plan.NetworkActionType
	}{
		{name: "overlay", network: fakeNetwork("", "default", ""),
			want: plan.NetworkOverrideDefault},
		{name: "remove", network: fakeNetwork("", "custom", "shared"),
			want: plan.NetworkRemoveConflict},
		{name: "no-op", network: fakeNetwork("", "default", "shared")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildNetworkPlan(
				fakeNetCaps(netname, true),
				[]compose.CollectedNetwork{tt.network},
				&fakeDockerInspector{})
			if tt.name == "no-op" {
				if len(got.Actions) > 0 {
					t.Fatalf("no-op should not be registerd")
				}
			} else {
				if len(got.Actions) != 1 {
					t.Fatalf("1 action should be in but %d in actions",
						len(got.Actions))
				}
				a := got.Actions[0]
				if a.Target.Key != tt.network.Ref.Key ||
					a.Target.Compose != tt.network.Ref.Compose {
					t.Fatalf("incorrect target: %+v", a.Target)
				}
				if a.ActionType != tt.want {
					t.Fatalf("expected action was %s but got %s",
						tt.want, a.ActionType)
				}
			}
		})
	}
}
