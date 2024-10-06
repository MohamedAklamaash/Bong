package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("./nids_monitor")

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(output))
}
