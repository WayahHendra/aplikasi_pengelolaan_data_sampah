package main

import (
	"fmt"
	cli_menus "trash-app/cli/cli-menus"
	gui_menus "trash-app/gui/gui-menus"
	"trash-app/utils"
)

func main() {
	var (
		isGui, input string // Variabel untuk menentukan mode GUI atau CLI dan input pengguna
		breakLoop    bool   // Penanda untuk keluar dari loop
	)

	utils.ClearConsole()

	// Meminta pengguna memilih mode GUI atau CLI
	fmt.Print("Tekan y untuk menggunakan GUI dan n untuk CLI: ")
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

			fmt.Scan(&input)
			value, err := utils.ValidateInput(input)

			if err != nil {
				fmt.Println("error:", err)
				return
			}

			cli_menus.HandleSelection(value, &breakLoop)
			if breakLoop {
				return // Keluar dari program jika breakLoop bernilai true
			}
		} else {
			fmt.Println("input tidak valid!")
		}
	}
}
