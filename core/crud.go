package core

import (
	"fmt"
	"sampah-app/utils"
)

func CreateData(wasteType string, recyclingMethod string, quantity float64, location string, status string) Waste {
	wasteType = GarbageTypes(wasteType)

	recyclingMethod = RecyclingMethods(recyclingMethod, wasteType)

	if utils.StrToLower(status) == "y" {
		status = "Sudah"
	} else if utils.StrToLower(status) == "n" {
		status = "Belum"
	} else {
		status = "nil"
	}

	return Waste{
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	}
}

func CreateManyData(data []Waste) []Waste {
	var (
		item, normalized Waste
		result           []Waste
	)

	for i := 0; i < len(data); i++ {
		item = data[i]
		normalized = CreateData(item.WasteType, item.RecyclingMethod, item.Quantity, item.Location, item.Status)
		result = append(result, normalized)
	}

	return result
}

func DeleteData(index int) {
	if index < 0 || index >= len(WasteData) {
		fmt.Println("index di luar batas")
	}

	// Geser data ke kiri
	for i := index; i < len(WasteData)-1; i++ {
		WasteData[i] = WasteData[i+1]
	}

	WasteData = WasteData[:len(WasteData)-1] // Potong slice 1 elemen terakhir
}
