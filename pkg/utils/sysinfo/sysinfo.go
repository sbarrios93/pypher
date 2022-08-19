package sysinfo

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type commonPaths struct {
	Currentwd string
	UserHome  string
}

var usr, _ = user.Current()

var CommonPaths = commonPaths{
	Currentwd: Getwd(),
	UserHome:  usr.HomeDir,
}

func PythonVersion() string {
	command := exec.Command("python", "-V")

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(out.String(), " ")[1]

}

func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory, got error %v", err)
	}
	return dir
}
