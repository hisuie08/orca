package docker

import (
	"fmt"
	"testing"
)

func TestDockerExecutor(t *testing.T) {
	testCases := []struct {
		desc   string
		allow  bool
		wantOp int
	}{
		{desc: "real", allow: true, wantOp: 4},
		{desc: "dry", allow: false, wantOp: 0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s_allow=%v", tC.desc, tC.allow), func(t *testing.T) {
			fake := newFakeExecutor(tC.allow).(*fakeExecutor)
			if _, e := fake.ComposeUp(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.ComposeDown(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.CreateVolume(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.CreateNetwork(""); e != nil {
				t.Fatal(e)
			}
			if len(fake.Ops) != tC.wantOp {
				t.Errorf("expected %d in Ops but got %d", len(fake.Ops), tC.wantOp)
			}
		})
	}
}
