package main

import (
	"aplikasi_pengelolaan_data_sampah/internal/validation"
	"fmt"
)

func main() {
	clear() // Clear console

	var input string

	// loop
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
			clear() // Clear console
		case 10:
			fmt.Println("Keluar dari program!")
			return // break loop
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
