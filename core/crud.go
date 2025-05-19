package core

import (
	"fmt"
	"trash-app/utils"
)

// CreateData membuat/memperbarui data sampah dengan normalisasi atribut tertentu.
func CreateData(wasteType string, recyclingMethod string, quantity float64, location string, status string) Waste {
	wasteType = GarbageTypes(wasteType) // Normalisasi jenis sampah

	recyclingMethod = RecyclingMethods(recyclingMethod, wasteType) // Normalisasi metode daur ulang berdasarkan jenis sampah

	// Normalisasi status (sudah atau belum didaur ulang)
	if utils.StrToLower(status) == "yes" || utils.StrToLower(status) == "ye" || utils.StrToLower(status) == "y" || utils.StrToLower(status) == "sudah" {
		if SwitchLanguage {
			status = "Done"
		} else {
			status = "Sudah"
		}
	} else if utils.StrToLower(status) == "no" || utils.StrToLower(status) == "n" || utils.StrToLower(status) == "belum" {
		if SwitchLanguage {
			status = "Not Done"
		} else {
			status = "Belum"
		}
	} else {
		status = "nil"
	}

	// Mengembalikan data sampah yang telah dinormalisasi
	return Waste{
		WasteType:       wasteType,
		RecyclingMethod: recyclingMethod,
		Quantity:        quantity,
		Location:        location,
		Status:          status,
	}
}

// CreateManyData membuat banyak data sampah sekaligus.
func CreateManyData(data []Waste) []Waste {
	var (
		item, normalized Waste   // Variabel untuk menyimpan data sampah yang dinormalisasi
		result           []Waste // Slice untuk menyimpan hasil normalisasi
	)

	// Iterasi setiap data sampah untuk dinormalisasi
	for i := 0; i < len(data); i++ {
		item = data[i]
		normalized = CreateData(item.WasteType, item.RecyclingMethod, item.Quantity, item.Location, item.Status)
		result = append(result, normalized) // Menambahkan data yang dinormalisasi ke dalam slice hasil
	}

	return result // Mengembalikan daftar data sampah yang telah dinormalisasi
}

// DeleteData menghapus data sampah berdasarkan indeks.
func DeleteData(index int) {
	// Validasi indeks agar tidak di luar batas
	if index < 0 || index >= len(WasteData) {
		if SwitchLanguage {
			fmt.Println("Index out of range. Please enter a valid index.")
		} else {
			fmt.Println("Indeks di luar jangkauan. Silakan masukkan indeks yang valid.")
		}
		return
	}

	// Geser elemen-elemen setelah indeks ke kiri
	for i := index; i < len(WasteData)-1; i++ {
		WasteData[i] = WasteData[i+1] // Geser elemen ke kiri
	}

	WasteData = WasteData[:len(WasteData)-1] // Potong slice untuk menghapus elemen terakhir
}
