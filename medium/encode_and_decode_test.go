package medium

import (
	"reflect"
	"testing"
)

func Test_encode_decode(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case #1",
			args: args{strs: []string{"lint", "code", "love", "you"}},
		},
		{
			name: "Test case #1",
			args: args{strs: []string{"we", "say", ":", "yes"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := encode(tt.args.strs)
			println(encoded)
			if got := decode(encoded); !reflect.DeepEqual(got, tt.args.strs) {
				t.Errorf("decode() = %v, want %v", got, tt.args.strs)
			}
		})
	}
}
