package internal

import "testing"

func Test_label(t *testing.T) {
	testString := "str_ptr"
	strptr := &testString

	tests := []struct {
		name string
		v    any
		want string
	}{
		// TODO: Add test cases.
		{name: "string", v: "value", want: "key=value"},
		{name: "bool", v: true, want: "key=true"},
		{name: "int", v: 1, want: "key=1"},
		{name: "ptr", v: *strptr, want: "key=str_ptr"},
		{name: "nil", v: nil, want: "key=<nil>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := label("key", tt.v)
			if got != tt.want {
				t.Errorf("label() = %v, want %v", got, tt.want)
			}
		})
	}
}
