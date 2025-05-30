package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// PressToContinue menunggu input dari pengguna sebelum melanjutkan.
func PressToContinue(text ...string) {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Print(text)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// ClearConsole membersihkan layar konsol.
func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
