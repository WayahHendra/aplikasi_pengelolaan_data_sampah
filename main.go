package main

import (
	"fmt"
	"sampah-app/cli/menus"
	"sampah-app/cli/utils"
	"sampah-app/validation"
)

func main() {
	var (
		isGui, input string
		breakLoop    bool
	)

	utils.ClearConsole()

	fmt.Print("Tekan y untuk menggunakan GUI dan n untuk CLI: ")
	fmt.Scan(&isGui)

	for {
		utils.ClearConsole()

		if utils.StrToLower(isGui) == "y" {
			// TODO: Implement GUI
		} else if utils.StrToLower(isGui) == "n" {
			menus.ShowTableMenu()

			fmt.Scan(&input)
			value, err := validation.ValidateInput(input)

			if err != nil {
				fmt.Println("error:", err)
				return
			}

			menus.HandleSelection(value, &breakLoop)
			if breakLoop {
				return
			}
		} else {
			fmt.Println("input tidak valid!")
		}
	}
}
