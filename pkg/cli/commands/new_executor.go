package commands

import (
	"fmt"
	"path"
	"runtime"
)

func RunNew(dirPath string) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	fmt.Printf("%q, %q, %q", filename, dir, dirPath)
}
