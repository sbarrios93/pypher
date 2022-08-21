//go:build !windows
// +build !windows

package paths

import (
	"reflect"
	"testing"
)

func TestRootDirectory(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want *ProjectPath
	}{
		{
			name: "Test root directory",
			args: args{dir: "/"},
			want: &ProjectPath{
				Path:   "/",
				Name:   "/",
				Parent: "/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsProjectPath(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
