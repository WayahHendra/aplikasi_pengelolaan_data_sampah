package main

import (
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // untuk Windows
	} else {
		cmd = exec.Command("clear") // untuk Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
