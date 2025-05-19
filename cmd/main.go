package main

import (
	"fmt"
	cli_menus "trash-app/cli/cli-menus"
	"trash-app/core"
	gui_menus "trash-app/gui/gui-menus"
	"trash-app/utils"
)

func main() {
	var (
		isGui, selectLanguage, menuSelection string // Variabel untuk menyimpan pilihan GUI, bahasa dan menu
		breakLoop                            bool   // Penanda untuk keluar dari loop
	)

	utils.ClearConsole()

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

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Print("Do you want to use GUI? (y/n): ")
	} else {
		fmt.Print("Apakah Anda ingin menggunakan GUI? (y/n): ")
	}
	fmt.Scan(&isGui)

	for {
		utils.ClearConsole()

		if utils.StrToLower(isGui) == "yes" || utils.StrToLower(isGui) == "ye" || utils.StrToLower(isGui) == "y" {
			// Jika pengguna memilih GUI
			gui_menus.RunGUI()
			return
		} else if utils.StrToLower(isGui) == "no" || utils.StrToLower(isGui) == "n" {
			// Jika pengguna memilih CLI
			breakLoop = false
			cli_menus.ShowTableMenu()

			fmt.Scan(&menuSelection)
			value, err := utils.ValidateInput(menuSelection)

			if err != nil {
				fmt.Println("error:", err)
				return
			}

			cli_menus.HandleSelection(value, &breakLoop)
			if breakLoop {
				return // Keluar dari program jika breakLoop bernilai true
			}
		} else {
			if core.SwitchLanguage {
				fmt.Println("Invalid input!")
			} else {
				fmt.Println("input tidak valid!")
			}
		}
	}
}
