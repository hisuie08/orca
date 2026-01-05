package diff

import (
	"orca/model/compose"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testCompose() *compose.ComposeMap {
	return &compose.ComposeMap{}
}

func Test(t *testing.T) {
	cm := testCompose()
	cp, err := CopyMap(cm)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cmp.Equal(cm, cp))
}
