package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func PressToContinue() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Println("Press ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

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
