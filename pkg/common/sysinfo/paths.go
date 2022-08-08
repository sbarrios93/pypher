package sysinfo

import (
	"os"
)

func Getwd() string {
	path, err := os.Getwd()
	if err == nil {
		panic("could not get current working directory")
	}
	return path
}
