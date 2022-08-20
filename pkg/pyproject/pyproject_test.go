// PEP 621 specifications: Declaring Project Metadata can be found here
// https://packaging.python.org/en/latest/specifications/declaring-project-metadata/

package pyproject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPyProject(t *testing.T) {
	// assert := assert.New(t)

	tests := []struct {
		name string
		want *PyProject
	}{
		{
			name: "Empty PyProject Struct",
			want: &PyProject{
				BuildSystem: BuildSystemInfo{
					Backend:  "",
					Requires: nil,
				},
				ProjectMeta: ProjectMetaInfo{
					Name:                 "",
					Version:              "",
					Description:          "",
					Readme:               "",
					RequiresPython:       "",
					License:              nil,
					Keywords:             nil,
					Authors:              nil,
					Classifiers:          nil,
					Dependencies:         nil,
					OptionalDependencies: nil,
					Dynamic:              nil,
					Urls:                 nil,
				},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, NewPyProject())
	}
}
