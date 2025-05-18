package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func ReadWaste() {
	utils.ClearConsole()

	var waste core.Waste

	if len(core.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Println("========================================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
	fmt.Println("||====================================================================================================||")
	for i := 0; i < len(core.WasteData); i++ { // Loop data
		waste = core.WasteData[i]
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("========================================================================================================")

	fmt.Println() // Spacing

	if core.TriggerShowData {
		utils.PressToContinue()
		utils.ClearConsole()
	}
}
