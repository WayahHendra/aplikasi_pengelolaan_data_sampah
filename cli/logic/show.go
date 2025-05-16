package logic

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/utils"
)

func ShowAllData() {
	utils.ClearConsole()

	var waste data.Waste

	if len(data.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Println("========================================================================================================")
	fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
	fmt.Println("||====================================================================================================||")
	for i := 0; i < len(data.WasteData); i++ { // Loop data
		waste = data.WasteData[i]
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("========================================================================================================")

	fmt.Println() // Spacing

	if data.TriggerShowData {
		utils.PressToContinue()
		utils.ClearConsole()
	}
}
