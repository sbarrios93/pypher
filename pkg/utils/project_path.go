package paths

import (
	"log"
	"path/filepath"
)

type ProjectPath struct {
	Path   string
	Name   string
	Parent string
}

func ToProjectPath(dir string) *ProjectPath {
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("could not resolve path specified. Got %s", dir)
	}

	cleanDir := filepath.Clean(dir)
	if err != nil {
		log.Fatalf("could not clean path specified. Got %s", cleanDir)
	}

	projectPath := &ProjectPath{
		Path:   cleanDir,
		Name:   filepath.Base(cleanDir),
		Parent: filepath.Dir(cleanDir),
	}

	return projectPath
}

func (p *ProjectPath) MkDir() {
	// TODO
}

func (p *ProjectPath) IsEmpty() {
	// TODO
}

func (p *ProjectPath) Exists() {
	// TODO
}
