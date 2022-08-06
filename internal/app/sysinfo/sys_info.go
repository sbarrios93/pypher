package sysinfo

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)


func PythonVersion() {
	command := exec.Command("python", "-V")

	command.Stdin = strings.NewReader("and old falcon")

    var out bytes.Buffer
    command.Stdout = &out

    err := command.Run()

    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("Python Version: %q\n", out.String())

}