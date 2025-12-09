package compose

import (
	"fmt"
	"orca/internal/yamlparser"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

type NetworkSpec struct {
	Name     string `yaml:"name"`
	External bool   `yaml:"external"`
	Driver   string `yaml:"driver"`
}
type NetStruct map[string]any

func (n NetStruct) FromNode(netNode *yaml.Node) *NetStruct {
	netStruct := NetStruct{}
	yamlparser.NodeToStruct(netNode, &netStruct)
	return &netStruct
}

// 
func OverlayNetwork(node *yaml.Node, name string, path string) {
	spec := NetworkSpec{
		Name:     name,
		External: true,
	}
	yamlparser.OverlayNode(node, "default", spec)
}

func NetworkProcess(data []byte) {
	root, _ := yamlparser.ParseYamlToNode(data)

	// networksノードを取得
	_, netNode := yamlparser.FindMapKey(root, "networks")
	// structに落とし込み
	netStruct := NetStruct{}.FromNode(netNode)

	spew.Dump(netStruct)
	// overlay 実行
	//TODO: オーバーレイの条件式と具体実装

	// 出力
	out, _ := yaml.Marshal(&root)
	fmt.Println(string(out)) //TODO: 出力の統合
}
