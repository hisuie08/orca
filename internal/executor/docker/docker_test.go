package docker

import (
	"fmt"
	"testing"
)

func TestDockerExecutor(t *testing.T) {
	testCases := []struct {
		desc       string
		allow      bool
		wantIssued int
		wantDone   int
	}{
		{desc: "real", allow: true, wantIssued: 4, wantDone: 4},
		{desc: "dry", allow: false, wantIssued: 4, wantDone: 0},
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
			if len(fake.Issued) != tC.wantIssued {
				t.Errorf("expected %d in Issued but got %d", len(fake.Issued), tC.wantIssued)
			}
			if len(fake.Done) != tC.wantDone {
				t.Errorf("expected %d in Done but got %d", len(fake.Done), tC.wantDone)
			}
		})
	}
}
