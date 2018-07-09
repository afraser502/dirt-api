package main

import (
	"fmt"
	"os/exec"
)

func checkTwistcliExists() {
	path, err := exec.LookPath("twistcli")
	if err != nil {
		fmt.Printf("Can't find 'twistcli' executable\n")
	} else {
		fmt.Printf("'twistcli' executable is in '%s'\n", path)
	}
}
