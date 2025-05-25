package core

import (
	"fmt"
	"trash-app/utils"
)

// SelectionSort mengurutkan data menggunakan algoritma Selection Sort.
func SelectionSort(sortField string, sortAscending bool) {
	for i := 0; i < len(WasteData)-1; i++ {
		var minOrMaxIdx int = i
		for j := i + 1; j < len(WasteData); j++ {
			// Tentukan elemen minimum/maksimum berdasarkan field yang dipilih
			switch sortField {
			case "jenis":
				if (sortAscending && WasteData[j].WasteType < WasteData[minOrMaxIdx].WasteType) || (!sortAscending && WasteData[j].WasteType > WasteData[minOrMaxIdx].WasteType) {
					minOrMaxIdx = j
				}
			case "metode":
				if (sortAscending && WasteData[j].RecyclingMethod < WasteData[minOrMaxIdx].RecyclingMethod) || (!sortAscending && WasteData[j].RecyclingMethod > WasteData[minOrMaxIdx].RecyclingMethod) {
					minOrMaxIdx = j
				}
			case "jumlah":
				if (sortAscending && WasteData[j].Quantity < WasteData[minOrMaxIdx].Quantity) || (!sortAscending && WasteData[j].Quantity > WasteData[minOrMaxIdx].Quantity) {
					minOrMaxIdx = j
				}
			case "lokasi":
				if (sortAscending && WasteData[j].Location < WasteData[minOrMaxIdx].Location) || (!sortAscending && WasteData[j].Location > WasteData[minOrMaxIdx].Location) {
					minOrMaxIdx = j
				}
			}
		}

		WasteData[i], WasteData[minOrMaxIdx] = WasteData[minOrMaxIdx], WasteData[i] // Tukar elemen
	}
}

// InsertionSort mengurutkan data menggunakan algoritma Insertion Sort.
func InsertionSort(sortField string, sortAscending bool) {
	for i := 1; i < len(WasteData); i++ {
		var key Waste = WasteData[i]
		j := i - 1

		// Geser elemen untuk menemukan posisi yang tepat
		for j >= 0 {
			shouldSwap := false
			switch sortField {
			case "jenis":
				shouldSwap = (sortAscending && WasteData[j].WasteType > key.WasteType) || (!sortAscending && WasteData[j].WasteType < key.WasteType)
			case "metode":
				shouldSwap = (sortAscending && WasteData[j].RecyclingMethod > key.RecyclingMethod) || (!sortAscending && WasteData[j].RecyclingMethod < key.RecyclingMethod)
			case "jumlah":
				shouldSwap = (sortAscending && WasteData[j].Quantity > key.Quantity) || (!sortAscending && WasteData[j].Quantity < key.Quantity)
			case "lokasi":
				shouldSwap = (sortAscending && WasteData[j].Location > key.Location) || (!sortAscending && WasteData[j].Location < key.Location)
			}

			if !shouldSwap {
				break
			}

			WasteData[j+1] = WasteData[j] // Geser elemen ke kanan
			j--
		}

		WasteData[j+1] = key // Sisipkan elemen pada posisi yang tepat
	}
}

// SetSortField menentukan field pengurutan berdasarkan input pengguna.
func SetSortField(output *string, input string) {
	switch utils.StrToLower(input) {
	case "1":
		*output = "jenis" // Field: WasteType
	case "2":
		*output = "metode" // Field: RecyclingMethod
	case "3":
		*output = "jumlah" // Field: Quantity
	case "4":
		*output = "lokasi" // Field: Location
	default:
		if SwitchLanguage {
			fmt.Println("Invalid choice. Use 1 for WasteType, 2 for RecyclingMethod, 3 for Quantity, and 4 for Location.")
		} else {
			fmt.Println("Pilihan tidak valid. Gunakan 1 untuk WasteType, 2 untuk RecyclingMethod, 3 untuk Quantity, dan 4 untuk Location.")
		}
		*output = ""
	}
}

// SetSortOrder menentukan urutan pengurutan (Ascending/Descending) berdasarkan input pengguna.
func SetSortOrder(output *bool, input string) {
	switch utils.StrToLower(input) {
	case "1":
		*output = true // Ascending
	case "2":
		*output = false // Descending
	default:
		if SwitchLanguage {
			fmt.Println("Invalid choice. Use 1 for Ascending, 2 for Descending.")
		} else {
			fmt.Println("Pilihan tidak valid. Gunakan 1 untuk Ascending, 2 untuk Descending.")
		}
	}
}

// SortByAlgorithm memilih algoritma pengurutan berdasarkan input pengguna.
func SortByAlgorithm(algorithmChoice string, sortField string, sortAscending bool) {
	switch algorithmChoice {
	case "1":
		SelectionSort(sortField, sortAscending) // Gunakan Selection Sort
	case "2":
		InsertionSort(sortField, sortAscending) // Gunakan Insertion Sort
	default:
		if SwitchLanguage {
			fmt.Println("Invalid choice. Use 1 for Selection Sort, 2 for Insertion Sort.")
		} else {
			fmt.Println("Pilihan tidak valid. Gunakan 1 untuk Selection Sort, 2 Insertion Sort.")
		}
	}
}
