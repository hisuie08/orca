package network

import (
	"orca/model/compose"
	"orca/model/plan"
	"testing"
)

func fakeNetworkCompose(c, k, n string) compose.ComposeMap {
	return compose.ComposeMap{c: &compose.ComposeSpec{
		Networks: compose.NetworksSection{
			k: &compose.NetworkSpec{
				Name:   n,
				Labels: map[string]string{"label": "value"}}},
	}}
}

func fakeNetworkPlan(s string, a plan.NetworkActionType, c, k string) plan.NetworkPlan {
	return plan.NetworkPlan{
		SharedName: s,
		Actions: []plan.NetworkAction{{
			ActionType: a,
			Target:     plan.NetworkRef{Compose: c, Key: k},
		}},
	}
}
func TestOverlayNetwork(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		composename string
		key         string
		action      plan.NetworkActionType
		want        compose.NetworksSection
	}{
		// TODO: Add test cases.
		{name: "override", action: plan.NetworkOverrideDefault},
		{name: "remove", action: plan.NetworkRemoveConflict},
	}
	for _, tt := range tests {
		netname := "shared"
		c := fakeNetworkCompose(tt.composename, tt.key, "")
		p := fakeNetworkPlan(netname, tt.action, tt.composename, tt.key)
		t.Run(tt.name, func(t *testing.T) {
			OverlayNetwork(c, p)
			switch tt.action {
			case plan.NetworkOverrideDefault:
				n := c[tt.composename].Networks[tt.key]
				if n.Name != netname || // is the name overriden?
					!n.External || // is it external?
					n.Labels["label"] != "value" { // is the label preserved?
					t.Fatalf("override failed %#v", *n)
				}
			case plan.NetworkRemoveConflict:
				if _, ok := c[tt.composename].Networks[tt.key]; ok {
					t.Fatalf("There are still networks to be deleted")
				}
			}

		})
	}
}

func TestMultipleCompose(t *testing.T) {
	c, k, n := "a", "default", ""
	composes := fakeNetworkCompose(c, k, n)
	composes["other"] = &compose.ComposeSpec{Networks: compose.NetworksSection{
		"default": &compose.NetworkSpec{Name: "othernetwork"},
	}}
	p := fakeNetworkPlan("shared", plan.NetworkOverrideDefault, c, k)
	OverlayNetwork(composes, p)
	if composes[c].Networks[k].Name != "shared" {
		t.Fatal("override failed")
	}
	if composes["other"].Networks[k].Name != "othernetwork" {
		t.Fatal("Unexpected compose was overwritten")
	}
}

func TestEmptyPlan(t *testing.T) {
	c, k, n := "a", "default", ""
	composes := fakeNetworkCompose(c, k, n)
	p := plan.NetworkPlan{SharedName: "shared", Actions: []plan.NetworkAction{}, Create: true}
	OverlayNetwork(composes, p)
	if composes[c].Networks[k].Name != "" {
		t.Fatal("Unexpected compose was overwritten")
	}
}
