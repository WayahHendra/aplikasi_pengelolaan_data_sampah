package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func SwitchLanguage() {
	utils.ClearConsole()

	var selectLanguage string // Variabel untuk menyimpan pilihan bahasa

	if core.SwitchLanguage {
		fmt.Println("Select language:")
		fmt.Println("[ 1] English")
		fmt.Println("[ 2] Indonesia")
		fmt.Println("[-1] Back to main menu")
		fmt.Print("Select language [1/2/-1]: ")
	} else {
		fmt.Println("Pilih bahasa:")
		fmt.Println("[ 1] English")
		fmt.Println("[ 2] Indonesia")
		fmt.Println("[-1] Kembali ke menu utama")
		fmt.Print("Pilih bahasa [1/2/-1]: ")
	}
	fmt.Scan(&selectLanguage)
	if selectLanguage == "1" {
		// Jika pengguna memilih bahasa Inggris
		core.SwitchLanguage = true
		fmt.Println("You have selected English")
	} else if selectLanguage == "2" {
		// Jika pengguna memilih bahasa Indonesia
		core.SwitchLanguage = false
		fmt.Println("Anda telah memilih bahasa Indonesia")
	} else if selectLanguage == "-1" {
		return // Kembali ke menu utama
	} else {
		fmt.Println("ID: Input tidak valid! Silakan pilih bahasa yang benar.")
		fmt.Println("EN: Invalid input! Please select a valid language.")
		return
	}
}
