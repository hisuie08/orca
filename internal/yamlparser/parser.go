package yamlparser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// YAML 読み取り
func ParseYamlToNode(data []byte) (*yaml.Node, error) {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil ||
		len(root.Content) != 1 {
		return nil, fmt.Errorf("yaml unmarshal error: %w", err)
	}
	return root.Content[0], nil
}

// キー抽出
// 
// root の mapping から特定 key を探す
func FindMapKey(root *yaml.Node, key string) (keyNode *yaml.Node, valueNode *yaml.Node) {
	if root.Kind != yaml.MappingNode {
		return nil, nil
	}
	for i := 0; i < len(root.Content); i += 2 {
		k := root.Content[i]
		if k.Value == key {
			return k, root.Content[i+1]
		}
	}
	return nil, nil
}

// Node → 任意 struct にパース
func NodeToStruct(node *yaml.Node, out any) error {
	b, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, out)
}


// Compose の再構築
func OverlayNode(node *yaml.Node, name string, newSpec any) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("volumes node must be a mapping")
	}

	var specNode yaml.Node
	b, _ := yaml.Marshal(newSpec)
	yaml.Unmarshal(b, &specNode)

	// Document → Mapping に展開
	if len(specNode.Content) == 1 {
		specNode = *specNode.Content[0]
	}

	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		if keyNode.Value == name {
			node.Content[i+1] = &specNode
			keyNode.LineComment = "orca modified"
			return nil
		}
	}

	return fmt.Errorf("not found: %s", name)
}
