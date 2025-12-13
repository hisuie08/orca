package compose

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

/*
通常のcompose.yml読み出しではなく `docker compose config` の出力を受け取るので
セクションは正規化される。ほぼ
*/
var testCompose = `
volumes:
  defaultvol:
    name: docker_defaultvol
  localvol:
    name: docker_localvol
    driver: local
    driver_opts:
      type: none
      o: bind
      device: /src/test
  namedvol:
    name: customname
  externalvol:
    external: true
  cache:
    driver: local
    driver_opts:
      type: tmpfs
networks:
  net:
    driver: bridge
  default:
    name: main
    external: true
`

func TestParseCompose(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test", []byte(testCompose), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ParseCompose(tt.data)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseCompose() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseCompose() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				spew.Dump(got)
			}
		})
	}
}
