package volume

import (
	"orca/internal/context"
	"orca/model/compose"
	"orca/model/config"
	"orca/model/plan"
	"path/filepath"
	"testing"
)

type fakeDockerInspector struct {
	Volume   string
	BindPath string
}

func (f *fakeDockerInspector) VolumeExists(name string) bool {
	return name == f.Volume
}
func (f *fakeDockerInspector) BindExists(dir string) bool {
	return dir == f.BindPath
}

// TestNilRoot 以外のテストで使うContext
func fakeVolCtx(root string, ensurePath bool) VolumePlanContext {
	ctx := context.New().WithConfig(&config.ResolvedConfig{
		Volume: config.ResolvedVolume{
			VolumeRoot: &root, EnsurePath: ensurePath}})
	return &ctx
}

// テスト用ボリュームビルダー
type mockVolume struct {
	Compose    string
	Key        string
	Driver     string
	DriverOpts map[string]string
	External   bool
	Labels     map[string]string
	Name       string
}

func (m *mockVolume) Build() compose.CollectedVolume {
	ref := compose.FromRef{Compose: m.Compose, Key: m.Key}
	spec := &compose.VolumeSpec{
		Driver:     m.Driver,
		DriverOpts: map[string]string{},
		External:   m.External,
		Labels:     map[string]string{},
		Name:       m.Name,
	}
	if m.DriverOpts != nil {
		spec.DriverOpts = m.DriverOpts
	}
	if m.Labels != nil {
		spec.Labels = m.Labels
	}
	return compose.CollectedVolume{Ref: ref, Spec: spec}
}

var di = &fakeDockerInspector{Volume: "exist_vol", BindPath: "/path/to/exist"}

var (
	/*caseExternal expects
	Type: external
	Quantity: 2
	usedBy: exist_vol=1,b_vol=2
	*/
	caseExternal = []compose.CollectedVolume{
		//VolumeExists → External
		(&mockVolume{Compose: "a", Key: "a", Name: "exist_vol"}).Build(),
		// b_volのうち最低1つがExternal 指定あり → External
		(&mockVolume{Compose: "b", Key: "b", Name: "b_vol", External: false}).Build(),
		(&mockVolume{Compose: "c", Key: "c", Name: "b_vol", External: true}).Build(),
	}
	/* caseShared expects
	Type: shared
	Quantity: 1
	usedBy: a_vol=2
	*/
	caseShared = []compose.CollectedVolume{
		//重複 かつ 全て !External → Shared
		(&mockVolume{Compose: "a", Key: "a", Name: "a_vol"}).Build(),
		(&mockVolume{Compose: "b", Key: "b", Name: "a_vol"}).Build(),
	}
	/* caseLocal expects
	Type: local
	Quantity: 3
	usedBy: a_vol=1,b_vol=1,c_vol=1
	*/
	caseLocal = []compose.CollectedVolume{
		// 単一箇所定義 かつ !External → Local
		(&mockVolume{Compose: "a", Key: "a", Name: "a_vol"}).Build(),
		(&mockVolume{Compose: "b", Key: "b", Name: "b_vol"}).Build(),
		(&mockVolume{Compose: "c", Key: "c", Name: "c_vol"}).Build(),
	}
)
var (
	/* caseSkip expects
	Quantity: 2
	usedBy: b_vol=1
	*/
	caseSkip = []compose.CollectedVolume{
		(&mockVolume{Compose: "a", Key: "a", Name: "a_vol"}).Build(),
		// スキップ対象 tmpfsはcompose-scoped volumeなのでc.cに影響しない
		(&mockVolume{Compose: "b", Key: "b", Name: "b_vol",
			Driver: "local", DriverOpts: map[string]string{"type": "tmpfs"}}).Build(),
		// スキップされない
		(&mockVolume{Compose: "c", Key: "c", Name: "b_vol", Driver: "local"}).Build(),
	}
)

func TestType(t *testing.T) {
	tests := []struct {
		name string
		cv   []compose.CollectedVolume
		want plan.VolumeType
	}{
		{name: "external", cv: caseExternal, want: plan.VolumeExternal},
		{name: "shared", cv: caseShared, want: plan.VolumeShared},
		{name: "local", cv: caseLocal, want: plan.VolumeLocal},
	}
	for _, tt := range tests {
		ctx := fakeVolCtx("testroot", true)
		t.Run("want "+tt.name, func(t *testing.T) {
			pl := buildVolumePlan(ctx, tt.cv, di)
			for _, p := range pl {
				if p.Type != tt.want {
					t.Fatalf("volumes should be %s; %#v", tt.want, p)
				}
			}
		})
	}
}
func TestAggregation(t *testing.T) {
	type Want struct {
		Quantity int
		UsedBy   map[string]int
	}
	tests := []struct {
		name string
		cv   []compose.CollectedVolume
		want Want
	}{ // 同名volume集約で減少あり cvExternal流用
		{name: "gatherd by name", cv: caseExternal,
			want: Want{Quantity: 2, UsedBy: map[string]int{
				"exist_vol": 1, "b_vol": 2}}},
		// 全て単一減少なし
		{name: "all local", cv: caseLocal, want: Want{
			Quantity: 3, UsedBy: map[string]int{
				"a_vol": 1, "b_vol": 1, "c_vol": 1,
			}}},

		// スキップありで減少あり
		{name: "has skip", cv: caseSkip, want: Want{Quantity: 2,
			UsedBy: map[string]int{"a_vol": 1, "b_vol": 1}}},
	}
	for _, tt := range tests {
		ctx := fakeVolCtx("testroot", true)
		di := &fakeDockerInspector{}
		t.Run(tt.name, func(t *testing.T) {
			pl := buildVolumePlan(ctx, tt.cv, di)
			wq := tt.want.Quantity
			lenpl := len(pl)
			t.Log(wq, lenpl)
			if lenpl != wq {
				t.Errorf("expected %d in plan but got %d", wq, lenpl)
			}
			for _, p := range pl {
				wu := tt.want.UsedBy[p.Name]
				lenu := len(p.UsedBy)
				if lenu != wu {
					t.Errorf("%s expected used by %d composes but got %d", p.Name, wu, lenu)
				}
			}
		})
	}
}
func TestNilRoot(t *testing.T) {
	// VolumeRoot nil → panic
	nilRoot := context.New().WithConfig(
		&config.ResolvedConfig{Volume: config.ResolvedVolume{
			VolumeRoot: nil, EnsurePath: true}})
	// VolumeRoot "" → panic
	es := ""
	emptyRoot := context.New().WithConfig(
		&config.ResolvedConfig{Volume: config.ResolvedVolume{
			VolumeRoot: &es, EnsurePath: true}})
	tests := []struct {
		name string
		ctx  VolumePlanContext
	}{
		{name: "nil root", ctx: &nilRoot}, {name: "empty root", ctx: &emptyRoot}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != "BuildVolumePlan called while volume management is disabled (VolumeRoot is nil)" {
					t.Errorf("test got unexpected error %v", err)
				}

			}()
			buildVolumePlan(tt.ctx, []compose.CollectedVolume{}, di)
		})
	}
}

func TestBind(t *testing.T) {
	fakeRoot := filepath.Join(t.TempDir(), "volRoot")
	cv := []compose.CollectedVolume{
		(&mockVolume{Compose: "a", Key: "a", Name: "a_vol",
			Driver: "local", DriverOpts: map[string]string{
				"type": "none", "o": "bind", "device": "/path/to/exist",
			}}).Build(),
		(&mockVolume{Compose: "b", Key: "b", Name: "b_vol",
			Driver: "local", DriverOpts: map[string]string{
				"type": "none", "o": "bind", "device": "/path/to/notexist",
			}}).Build(),
		(&mockVolume{Compose: "c", Key: "c", Name: "c_vol",
			Driver: "local"}).Build(),
	}
	want := map[string]struct {
		path      string
		needMkdir bool
	}{
		"a_vol": {path: "/path/to/exist", needMkdir: false},
		"b_vol": {path: "/path/to/notexist", needMkdir: true},
		// orca set default volume root <VolumeRoot>/<volume_name>
		"c_vol": {path: filepath.Join(fakeRoot, "c_vol"), needMkdir: true},
	}
	ctx := fakeVolCtx(fakeRoot, true)
	pl := buildVolumePlan(ctx, cv, di)
	for _, p := range pl {
		w := want[p.Name]
		if p.BindPath != w.path {
			t.Errorf("expected path %s but got %s", w.path, p.BindPath)
		}
		if p.NeedMkdir != w.needMkdir {
			t.Errorf("expected needMkdir %t but got %t", w.needMkdir, p.NeedMkdir)
		}
	}
}
