package sysinfo

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func PythonVersion() {
	command := exec.Command("python", "-V")

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Python Version: %q", out.String())

}

func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory, got error %v", err)
	}
	return dir
}
