package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// ShowStatistics menampilkan statistik jumlah sampah yang dikumpulkan dan jumlah yang berhasil didaur ulang.
func ShowStatistics() {
	utils.ClearConsole()

	var totalWaste, recycledWaste float64

	// Iterasi melalui data sampah untuk menghitung statistik
	for i := 0; i < len(core.WasteData); i++ {
		var waste core.Waste = core.WasteData[i]
		totalWaste += waste.Quantity
		if waste.Status == "Sudah" || waste.Status == "Complete" {
			recycledWaste += waste.Quantity
		}
	}

	// Tampilkan statistik
	fmt.Println("================================================")
	fmt.Println("||              Waste Statistics              ||")
	fmt.Println("================================================")
	if core.SwitchLanguage {
		fmt.Printf("|| %-30s|| %.2f kg ||\n", "Total waste collected", totalWaste)
		fmt.Printf("|| %-30s|| %.2f kg ||\n", "Total waste recycled", recycledWaste)
		fmt.Printf("|| %-30s|| %.2f%%    ||\n", "Recycling rate", (recycledWaste/totalWaste)*100)
	} else {
		fmt.Printf("|| %-30s|| %.2f kg ||\n", "Total sampah terkumpul", totalWaste)
		fmt.Printf("|| %-30s|| %.2f kg ||\n", "Total sampah didaur ulang", recycledWaste)
		fmt.Printf("|| %-30s|| %.2f%%    ||\n", "Persentase daur ulang", (recycledWaste/totalWaste)*100)
	}
	fmt.Println("================================================")

	// Tunggu input pengguna sebelum melanjutkan
	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to return to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
	}
	utils.ClearConsole()
}
