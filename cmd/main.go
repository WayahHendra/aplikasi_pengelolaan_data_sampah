package main

import (
	"aplikasi_pengelolaan_data_sampah/internal/validation"
	"fmt"
)

func main() {
	clearConsole() // Clear console
	// loadingScreen() // Loading screen :)

	var input, isGui string

	fmt.Print("Tekan y untuk menggunakan GUI dan n untuk CLI: ")
	fmt.Scan(&isGui)

	if strToLower(isGui) == "y" {
		// TODO:
		fmt.Println("GUI masih kosong maseh!")

		return
	} else if strToLower(isGui) == "n" {
		clearConsole() // Clear console

		// Loop
		for {
			showMenu()

			fmt.Scan(&input)
			value, err := validation.ValidateInput(input)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			switch value {
			case 1:
				addData()
			case 2:
				updateData()
			case 3:
				deleteData()
			case 4:
				recordProcess()
			case 5:
				searchData()
			case 6:
				sortData()
			case 7:
				showStatistics()
			case 8:
				showAllData()
			case 9:
				clearConsole() // Clear console
			case 10:
				fmt.Println("Keluar dari program!")
				return // Break loop
			default:
				fmt.Println("Pilihan tidak valid, coba lagi.")
			}
		}
	} else {
		fmt.Println("Input tidak valid!")
	}
}
