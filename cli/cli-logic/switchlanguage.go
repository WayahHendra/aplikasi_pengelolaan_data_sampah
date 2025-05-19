package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func SwitchLanguage() {
	var selectLanguage string // Variabel untuk menyimpan pilihan bahasa

	fmt.Print("Pilih bahasa (id/en): ")
	fmt.Scan(&selectLanguage)
	if utils.StrToLower(selectLanguage) == "en" {
		// Jika pengguna memilih bahasa Inggris
		core.SwitchLanguage = true
		fmt.Println("You have selected English")
	} else if utils.StrToLower(selectLanguage) == "id" {
		// Jika pengguna memilih bahasa Indonesia
		core.SwitchLanguage = false
		fmt.Println("Anda telah memilih bahasa Indonesia")
	} else {
		fmt.Println("ID: Input tidak valid! Silakan pilih bahasa yang benar.")
		fmt.Println("EN: Invalid input! Please select a valid language.")
		return
	}
}
