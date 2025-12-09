package yamlparser_test

import (
	"fmt"
	"orca/internal/compose"
	"orca/internal/yamlparser"
	"orca/testdata"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

var data string = testdata.TestDataCompose

func TestParseYaml(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		{"test1", []byte(data), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := yamlparser.ParseYamlToNode(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseYaml() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseYaml() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				fmt.Printf("%#v", got)
			}
		})
	}
}

func TestFindMapKey(t *testing.T) {
	node, _ := yamlparser.ParseYamlToNode([]byte(data))
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *yaml.Node
		key  string
	}{
		// TODO: Add test cases.
		{"test", node, "volumes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, volNode := yamlparser.FindMapKey(node, "volumes")
			if volNode == nil {
				fmt.Println("no volumes section")
				return
			}
			volMap := *compose.VolStruct{}.FromNode(volNode)
			spew.Dump(volMap)

		})
	}
}
