package sysinfo

import (
	"bytes"
	"fmt"
	"log"
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