package getall

import (
	"errors"
	"orca/errs"
	"orca/internal/capability"
	"orca/internal/errdef"
	"orca/model/compose"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type fakeDockerInspector struct {
	ComposeMap map[string]*compose.ComposeSpec
	ErrMap     map[string]error
}

func (f *fakeDockerInspector) Compose(dir string) (*compose.ComposeSpec, error) {
	if err, ok := f.ErrMap[dir]; ok {
		return nil, err
	}
	if c, ok := f.ComposeMap[dir]; ok {
		return c, nil
	}
	return nil, errs.ErrComposeNotFound
}

type fakeFSInspector struct {
	DirsMap map[string][]string
	Err     error
}

func (f *fakeFSInspector) Dirs(path string) ([]string, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.DirsMap[path], nil
}

var fakeRoot = "/root"

func TestGetALLCompose_Success(t *testing.T) {
	testCases := []struct {
		name     string
		dirs     []string
		composes map[string]*compose.ComposeSpec
		want     compose.ComposeMap
	}{
		{name: "All found", dirs: []string{
			fakeRoot + "/a", fakeRoot + "/b"},
			composes: map[string]*compose.ComposeSpec{
				fakeRoot + "/a": {}, fakeRoot + "/b": {}},
			want: compose.ComposeMap{"a": {}, "b": {}}},
		{name: "Skip not found", dirs: []string{
			fakeRoot + "/a", fakeRoot + "/b"},
			composes: map[string]*compose.ComposeSpec{fakeRoot + "/b": {}},
			want:     compose.ComposeMap{"b": {}}},
		{name: "Empty dirs", dirs: []string{},
			composes: map[string]*compose.ComposeSpec{},
			want:     compose.ComposeMap{}},
		{name: "Compose with contents", dirs: []string{fakeRoot + "/a"},
			composes: map[string]*compose.ComposeSpec{fakeRoot + "/a": &compose.ComposeSpec{
				Volumes: compose.VolumesSection{
					"a_vol": &compose.VolumeSpec{Name: "a_vol"}},
				Networks: compose.NetworksSection{
					"default": &compose.NetworkSpec{Name: "a_network"}}},
			}, want: compose.ComposeMap{"a": &compose.ComposeSpec{
				Volumes: compose.VolumesSection{
					"a_vol": &compose.VolumeSpec{Name: "a_vol"}},
				Networks: compose.NetworksSection{
					"default": &compose.NetworkSpec{Name: "a_network"}}}}},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			caps := capability.New().WithRoot(fakeRoot)
			fs := &fakeFSInspector{DirsMap: map[string][]string{
				fakeRoot: tt.dirs}}
			docker := &fakeDockerInspector{ComposeMap: tt.composes}
			result, err := getAllCompose(&caps, docker, fs)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			diff := cmp.Diff(tt.want, result, cmpopts.EquateEmpty(),
				cmpopts.IgnoreUnexported())
			if len(diff) != 0 {
				t.Errorf("%v", diff)
			}
		})
	}
}

func TestGetAllCompose_HasError(t *testing.T) {
	testCases := []struct {
		name string
		dErr error
		fErr error
	}{
		{name: "fs error", fErr: &errs.TestError{Err: errors.New("fs error")}},
		{name: "docker error", dErr: &errs.TestError{
			Err: errors.New("docker down")}}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			caps := capability.New().WithRoot(fakeRoot)
			fs := &fakeFSInspector{DirsMap: map[string][]string{
				fakeRoot: {fakeRoot + "/a"}}, Err: tt.fErr}
			docker := &fakeDockerInspector{
				ErrMap: map[string]error{fakeRoot + "/a": tt.dErr}}
			_, err := getAllCompose(&caps, docker, fs)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !errors.Is(err, errdef.TestErr) {
				t.Fatalf("unexpected error type: %v", err)
			}
		})
	}
}
