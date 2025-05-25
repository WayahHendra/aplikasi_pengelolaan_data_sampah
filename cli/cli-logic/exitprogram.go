package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func ExitProgram(breakLoop *bool) {
	utils.ClearConsole()

	var exitConfirmation string

	if core.SwitchLanguage {
		fmt.Println("Do you want to exit the program?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		fmt.Print("Select option [1/2]: ")
	} else {
		fmt.Println("Apakah Anda ingin keluar dari program?")
		fmt.Println("[1] Ya")
		fmt.Println("[2] Tidak")
		fmt.Print("Pilih opsi [1/2]: ")
	}

	fmt.Scan(&exitConfirmation)
	utils.ClearConsole()

	switch exitConfirmation {
	case "1":
		if core.SwitchLanguage {
			fmt.Println("Exiting the program...")
			utils.PressToContinue("Press Enter to continue...")
		} else {
			fmt.Println("Keluar dari program...")
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}
		*breakLoop = true // Set breakLoop menjadi true untuk keluar dari loop
		utils.ClearConsole()
	case "2":
		*breakLoop = false // Tetap di dalam loop
	default:
		if core.SwitchLanguage {
			fmt.Println("Invalid input! Please try again.")
			utils.PressToContinue("Press Enter to return to exit confirmation...")
		} else {
			fmt.Println("Input tidak valid! Silahkan coba lagi.")
			utils.PressToContinue("Tekan Enter untuk kembali ke konfirmasi keluar...")
		}
		ExitProgram(breakLoop) // Panggil kembali untuk meminta input ulang
	}
}
