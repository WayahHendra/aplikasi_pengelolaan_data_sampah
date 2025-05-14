package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func pressToContinue() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Println("Press ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Untuk Windows
	} else {
		cmd = exec.Command("clear") // Untuk Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func strToLower(s string) string {
	runes := []rune(s) // Ubah string ke slice rune agar mendukung karakter Unicode
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 'a' - 'A' // Ubah ke huruf kecil
		}
	}
	return string(runes) // Ubah kembali ke string
}

func strToUpper(s string) string {
	runes := []rune(s) // Ubah string ke slice rune agar mendukung karakter Unicode
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] -= 'a' - 'A' // Ubah ke huruf kapital
		}
	}
	return string(runes) // Ubah kembali ke string
}

func loadingScreen() { // Source from https://go.dev/play/
	const col = 30
	bar := fmt.Sprintf("Loading! [%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
		clearConsole()
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
	fmt.Println()
}
