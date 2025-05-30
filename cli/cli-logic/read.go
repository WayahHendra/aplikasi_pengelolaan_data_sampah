package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// ReadWaste menampilkan data sampah yang tersimpan.
func ReadWaste(date string) {
	utils.ClearConsole()

	var waste core.Waste // Variabel sementara untuk menyimpan data sampah

	// Periksa apakah ada data sampah yang tersedia
	if len(core.WasteData) == 0 {
		if core.SwitchLanguage {
			fmt.Println("No waste data available now, please add or reload waste data.")
		} else {
			fmt.Println("Belum ada data sampah yang tersedia saat ini, silahkan tambahkan atau muat ulang data sampah.")
		}

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to return to the main menu...")
		} else {
			utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
		}
		utils.ClearConsole()

		return
	}

	if date != "" { // Jika ada tanggal yang diberikan, tampilkan data sampah berdasarkan tanggal tersebut
		if core.SwitchLanguage {
			fmt.Printf("Waste data on date %s:\n", date)
		} else {
			fmt.Printf("Data sampah pada tanggal %s:\n", date)
		}
	}

	if core.SwitchLanguage {
		fmt.Println("============================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-12s ||\n", "No", "Type", "Recycling Method", "Quantity", "Location", "Status")
		fmt.Println("||========================================================================================================||")
	} else {
		fmt.Println("============================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-12s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
		fmt.Println("||========================================================================================================||")
	}
	// Iterasi untuk menampilkan setiap data sampah
	for i := 0; i < len(core.WasteData); i++ {
		waste = core.WasteData[i] // Ambil data sampah
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-12s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("============================================================================================================")

	fmt.Println()

	// Jika trigger aktif, tunggu input untuk melanjutkan dan bersihkan layar
	if core.TriggerShowData {
		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to return to the main menu...")
		} else {
			utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
		}

		utils.ClearConsole()
	}
}
