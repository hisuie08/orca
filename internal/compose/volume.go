package compose

import (
	"fmt"
	"orca/internal/yamlparser"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)


//	Node → Struct 変換で使用（判定用）
type VolumeSpec struct {
	Driver     string            `yaml:"driver,omitempty"`
	DriverOpts map[string]string `yaml:"driver_opts,omitempty"`
	External   bool              `yaml:"external,omitempty"`
	Labels     map[string]string `yaml:"labels,omitempty"`
	Name       string            `yaml:"name,omitempty"`
}
type VolStruct map[string]VolumeSpec

// ローカルバインドボリュームであればパスを返す
func (v VolumeSpec) AssertLocalBind() (string, bool) {
	isLocalDriver := v.Driver == "local" || v.Driver == ""
	isTypeNone := v.DriverOpts["type"] == "none"
	isBind := v.DriverOpts["o"] == "bind"
	device := v.DriverOpts["device"]
	conclusion := isLocalDriver && isTypeNone && isBind
	return device, conclusion

}

// local+bind+deviceが存在しないケース
func (v VolumeSpec) NeedsOrcaCreate() bool {
	device, isLocal := v.AssertLocalBind()
	if !isLocal {
		return false
	}
	// device が存在しない判定
	if _, err := os.Stat(device); os.IsNotExist(err) {
		return true
	}
	return false
}

func (v VolStruct) FromNode(volNode *yaml.Node) *VolStruct {
	volStruct := VolStruct{}
	yamlparser.NodeToStruct(volNode, &volStruct)
	return &volStruct
}

// name のボリュームを path にバインドするローカルボリュームに置き換え
func OverlayLocalVolume(node *yaml.Node, name string, path string) {
	spec := VolumeSpec{
		Driver: "local",
		DriverOpts: map[string]string{
			"o":      "bind",
			"type":   "none",
			"device": path + "/" + name,
		},
	}
	yamlparser.OverlayNode(node, name, spec)
}

func VolumeProcess(data []byte) {
	root, _ := yamlparser.ParseYamlToNode(data)

	// volumesノードを取得
	_, volNode := yamlparser.FindMapKey(root, "volumes")
	// structに落とし込み
	volStruct := VolStruct{}.FromNode(volNode)
	spew.Dump(volStruct)
	// overlay 実行
	//TODO: オーバーレイの条件式と具体実装

	// 出力
	out, _ := yaml.Marshal(&root)
	fmt.Println(string(out)) //TODO: 出力の統合
}
