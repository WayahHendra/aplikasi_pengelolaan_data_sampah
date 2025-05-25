package main

import (
	"fmt"
	cli_menus "trash-app/cli/cli-menus"
	"trash-app/core"
	"trash-app/gui"
	"trash-app/utils"
)

func main() {
	var (
		menuSelection         int    // Variabel untuk menyimpan pilihan menu
		isGui, selectLanguage string // Variabel untuk menyimpan pilihan GUI dan bahasa
		breakLoop             bool   // Penanda untuk keluar dari loop
	)

	utils.ClearConsole()

	fmt.Println("Pilih bahasa:")
	fmt.Println("[1] English")
	fmt.Println("[2] Indonesia")
	fmt.Print("Pilih bahasa [1/2]: ")
	fmt.Scan(&selectLanguage)
	if selectLanguage == "1" {
		// Jika pengguna memilih bahasa Inggris
		core.SwitchLanguage = true
		fmt.Println("You have selected English")
	} else if selectLanguage == "2" {
		// Jika pengguna memilih bahasa Indonesia
		core.SwitchLanguage = false
		fmt.Println("Anda telah memilih bahasa Indonesia")
	} else {
		fmt.Println("ID: Input tidak valid! Silakan pilih bahasa yang benar.")
		fmt.Println("EN: Invalid input! Please select a valid language.")
		return
	}

	utils.ClearConsole()

	if core.SwitchLanguage {
		fmt.Println("Do you want to use GUI?")
		fmt.Println("[1] Yes [not implemented yet!]")
		fmt.Println("[2] No")
		fmt.Print("Select mode [1/2]: ")
	} else {
		fmt.Println("Apakah Anda ingin menggunakan GUI?")
		fmt.Println("[1] Ya [belum diimplementasikan!]")
		fmt.Println("[2] Tidak")
		fmt.Print("Pilih mode [1/2]: ")
	}
	fmt.Scan(&isGui)

	for {
		utils.ClearConsole()

		if isGui == "1" {
			// Jika pengguna memilih GUI
			gui.RunGUI()
			return
		} else if isGui == "2" {
			// Jika pengguna memilih CLI
			breakLoop = false
			cli_menus.ShowTableMenu()

			fmt.Scan(&menuSelection)

			cli_menus.HandleSelection(menuSelection, &breakLoop)
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
