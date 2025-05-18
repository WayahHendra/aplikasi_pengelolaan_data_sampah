package core

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveWasteToFile menyimpan data sampah ke dalam file JSON.
func SaveWasteToFile(filename string, data []Waste) error {
	// Membuat file baru untuk menyimpan data
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w", err) // Mengembalikan error jika gagal membuat file
	}
	defer file.Close() // Menutup file setelah selesai

	// Membuat encoder JSON dengan indentasi
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	// Menyimpan data ke file dalam format JSON
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("%w", err) // Mengembalikan error jika encoding gagal
	}

	return nil // Berhasil menyimpan data
}

// LoadWasteFromFile memuat data sampah dari file JSON.
func LoadWasteFromFile(filename string) ([]Waste, error) {
	var rawData []Waste

	// Membuka file untuk membaca data
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%w", err) // Mengembalikan error jika gagal membuka file
	}
	defer file.Close() // Menutup file setelah selesai

	// Membuat decoder JSON untuk membaca data
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, fmt.Errorf("%w", err) // Mengembalikan error jika decoding gagal
	}

	return CreateManyData(rawData), nil // Mengembalikan data yang telah dinormalisasi
}
