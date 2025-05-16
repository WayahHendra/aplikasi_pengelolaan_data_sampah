package logic

import (
	"fmt"
	"sampah-app/cli/data"
	"sampah-app/cli/utils"
)

func SortData() {
	utils.ClearConsole()

	var (
		sortFieldChoice string // Input user (1/2/3)
		sortOrderChoice string // Input user (1/2)
		sortField       string // Output: "jenis", "metode", "jumlah"
		sortAscending   bool   // Output: true (ascending), false (descending)
	)

	fmt.Print("Urutkan berdasarkan (1) Jenis / (2) Metode Daur Ulang / (3) Jumlah? / (4) Lokasi: ")
	fmt.Scan(&sortFieldChoice)
	SetSortField(&sortField, sortFieldChoice)

	fmt.Print("Urutkan berdasarkan (1) Ascending / (2) Descending?: ")
	fmt.Scan(&sortOrderChoice)
	SetSortOrder(&sortAscending, sortOrderChoice)

	for i := 0; i < len(data.WasteData)-1; i++ {
		for j := 0; j < len(data.WasteData)-i-1; j++ {
			shouldSwap := false

			switch sortField {
			case "jenis":
				if sortAscending { // ascending
					shouldSwap = data.WasteData[j].WasteType > data.WasteData[j+1].WasteType
				} else { // descending
					shouldSwap = data.WasteData[j].WasteType < data.WasteData[j+1].WasteType
				}
			case "metode":
				if sortAscending { // ascending
					shouldSwap = data.WasteData[j].RecyclingMethod > data.WasteData[j+1].RecyclingMethod
				} else { // descending
					shouldSwap = data.WasteData[j].RecyclingMethod < data.WasteData[j+1].RecyclingMethod
				}
			case "jumlah":
				if sortAscending { // ascending
					shouldSwap = data.WasteData[j].Quantity > data.WasteData[j+1].Quantity
				} else { // descending
					shouldSwap = data.WasteData[j].Quantity < data.WasteData[j+1].Quantity
				}
			case "lokasi":
				if sortAscending { // ascending
					shouldSwap = data.WasteData[j].Location > data.WasteData[j+1].Location
				} else { // descending
					shouldSwap = data.WasteData[j].Location < data.WasteData[j+1].Location
				}
			default:
				fmt.Println("Pilihan pengurutan tidak dikenali.")
				return
			}

			if shouldSwap {
				data.WasteData[j], data.WasteData[j+1] = data.WasteData[j+1], data.WasteData[j] // Tukar elemen
			}
		}
	}

	switch sortField { // Pesan sukses
	case "jenis":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Jenis (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Jenis (Z-A).")
		}
	case "metode":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Metode Daur Ulang (Z-A).")
		}
	case "jumlah":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terkecil ke terbesar).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Jumlah (terbesar ke terkecil).")
		}
	case "lokasi":
		if sortAscending {
			fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (A-Z).")
		} else {
			fmt.Println("Data berhasil diurutkan berdasarkan Lokasi (Z-A).")
		}
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}

func SetSortField(output *string, input string) {
	switch utils.StrToLower(input) {
	case "1":
		*output = "jenis" // WasteType
	case "2":
		*output = "metode" // RecyclingMethod
	case "3":
		*output = "jumlah" // Quantity
	case "4":
		*output = "lokasi"
	default:
		fmt.Println("Pilihan pengurutan tidak valid. Gunakan 1, 2, 3, atau 4.")
		*output = ""
	}
}

func SetSortOrder(output *bool, input string) {
	switch utils.StrToLower(input) {
	case "1":
		*output = true // Ascending
	case "2":
		*output = false // Descending
	default:
		fmt.Println("Pilihan tidak valid. Gunakan 1 untuk ascending, 2 untuk descending.")
	}
}
