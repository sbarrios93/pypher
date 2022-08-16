package paths

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ProjectPath struct {
	Path   string
	Name   string
	Parent string
}

func AsProjectPath(dir string) *ProjectPath {
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

func (p *ProjectPath) MkDirAll() {
	err := os.MkdirAll(p.Path, os.ModePerm)
	if err != nil {
		log.Fatalf("could not make directories for path specified on %s. Got error %v", p.Path, err)
	}
}

func (p *ProjectPath) IsEmpty() bool {
	if p.Exists() {
		dirs, err := ioutil.ReadDir(p.Path)
		if err != nil {
			log.Fatalf("can't check if dir %s is empty", p.Path)
		}
		if len(dirs) == 0 {
			return true
		} else {
			return false
		}
	} else {
		log.Fatalf("cant check if dir %s is empty, because the dir does not exist", p.Path)
	}
	return false
}

func (p *ProjectPath) Exists() bool {
	_, err := os.Stat(p.Path)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
