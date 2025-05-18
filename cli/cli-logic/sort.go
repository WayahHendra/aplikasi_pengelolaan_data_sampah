package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

func SortWaste() {
	utils.ClearConsole()

	var (
		sortFieldChoice string // Input user (1/2/3)
		sortOrderChoice string // Input user (1/2)
		sortField       string // Output: "jenis", "metode", "jumlah"
		sortAscending   bool   // Output: true (ascending), false (descending)
	)

	if len(core.WasteData) == 0 { // Panjang data sampah == 0
		fmt.Println("Belum ada data sampah yang tersedia.")
		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Println("Urutkan data berdasarkan kolom apa?")
	fmt.Println("1. Jenis Sampah")
	fmt.Println("2. Metode Daur Ulang")
	fmt.Println("3. Jumlah Sampah (kg)")
	fmt.Println("4. Lokasi Pengumpulan")
	fmt.Print("Masukkan pilihan (1-4): ")
	fmt.Scan(&sortFieldChoice)
	core.SetSortField(&sortField, sortFieldChoice)

	fmt.Println() // Spacing

	fmt.Println("Urutan pengurutan:")
	fmt.Println("1. Naik (Ascending)")
	fmt.Println("2. Turun (Descending)")
	fmt.Print("Masukkan pilihan (1-2): ")
	fmt.Scan(&sortOrderChoice)
	core.SetSortOrder(&sortAscending, sortOrderChoice)

	fmt.Println() // Spacing

	for i := 0; i < len(core.WasteData)-1; i++ {
		for j := 0; j < len(core.WasteData)-i-1; j++ {
			shouldSwap := false

			switch sortField {
			case "jenis":
				if sortAscending { // ascending
					shouldSwap = core.WasteData[j].WasteType > core.WasteData[j+1].WasteType
					fmt.Println("Data berhasil diurutkan berdasarkan Jenis (A-Z).")
				} else { // descending
					shouldSwap = core.WasteData[j].WasteType < core.WasteData[j+1].WasteType
					fmt.Println("Data berhasil diurutkan berdasarkan Jenis (Z-A).")
				}
			case "metode":
				if sortAscending { // ascending
					shouldSwap = core.WasteData[j].RecyclingMethod > core.WasteData[j+1].RecyclingMethod
					fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (A-Z).")
				} else { // descending
					shouldSwap = core.WasteData[j].RecyclingMethod < core.WasteData[j+1].RecyclingMethod
					fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (Z-A).")
				}
			case "jumlah":
				if sortAscending { // ascending
					shouldSwap = core.WasteData[j].Quantity > core.WasteData[j+1].Quantity
					fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terkecil ke terbesar).")
				} else { // descending
					shouldSwap = core.WasteData[j].Quantity < core.WasteData[j+1].Quantity
					fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terbesar ke terkecil).")
				}
			case "lokasi":
				if sortAscending { // ascending
					shouldSwap = core.WasteData[j].Location > core.WasteData[j+1].Location
					fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (A-Z).")
				} else { // descending
					shouldSwap = core.WasteData[j].Location < core.WasteData[j+1].Location
					fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (Z-A).")
				}
			default:
				fmt.Println("Pilihan pengurutan tidak dikenali.")
				return
			}

			if shouldSwap {
				core.WasteData[j], core.WasteData[j+1] = core.WasteData[j+1], core.WasteData[j] // Tukar elemen
			}
		}
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
