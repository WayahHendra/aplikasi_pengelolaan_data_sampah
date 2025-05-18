package core

import (
	"fmt"
	"trash-app/utils"
)

func SelectionSort(sortField string, sortAscending bool) {
	for i := 0; i < len(WasteData)-1; i++ {
		minOrMaxIdx := i
		for j := i + 1; j < len(WasteData); j++ {
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
		WasteData[i], WasteData[minOrMaxIdx] = WasteData[minOrMaxIdx], WasteData[i]
	}
}

func InsertionSort(sortField string, sortAscending bool) {
	for i := 1; i < len(WasteData); i++ {
		key := WasteData[i]
		j := i - 1

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

			WasteData[j+1] = WasteData[j]
			j--
		}
		WasteData[j+1] = key
	}
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
		fmt.Println("Pilihan tidak valid. Gunakan 1 untuk Ascending, 2 untuk Descending.")
	}
}

func SortByAlgorithm(algorithmChoice string, sortField string, sortAscending bool) {
	switch algorithmChoice {
	case "1":
		SelectionSort(sortField, sortAscending)
	case "2":
		InsertionSort(sortField, sortAscending)
	default:
		fmt.Println("Pilihan tidak valid. Gunakan 1 untuk Selection Sort, 2 Insertion Sort.")
	}
}
