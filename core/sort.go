package core

import (
	"fmt"
	"sampah-app/utils"
)

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
