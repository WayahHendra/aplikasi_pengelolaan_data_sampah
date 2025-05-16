package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sampah-app/cli/utils"
)

func SaveWasteData(data []Waste) {
	utils.ClearConsole()

	var filename, confirmSave string

	fmt.Print("Masukkan nama file untuk menyimpan data: ")
	fmt.Scan(&filename)

	fmt.Print("Yakin ingin save file? (y/n): ")
	fmt.Scan(&confirmSave)

	if utils.StrToLower(confirmSave) == "y" {
		filename = utils.EnsureJSONExtension(filename)

		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Gagal membuat file:", err)

			fmt.Println() // Spacing

			utils.PressToContinue()
			utils.ClearConsole()

			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(data); err != nil {
			fmt.Println("Gagal menyimpan data ke file:", err)
		} else {
			fmt.Println("Data berhasil disimpan ke", filename)
		}
	} else if utils.StrToLower(confirmSave) == "n" {
		fmt.Println("Gagal menyimpan data!")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	} else {
		fmt.Println("Input tidak valid!")

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}

func LoadWasteData(waste *[]Waste) {
	utils.ClearConsole()

	var (
		filename string
		data     []Waste
	)

	fmt.Print("Masukkan nama file untuk load data: ")
	fmt.Scan(&filename)

	filename = utils.EnsureJSONExtension(filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Gagal membaca data dari file:", err)

		fmt.Println() // Spacing

		utils.PressToContinue()
		utils.ClearConsole()

		return
	}

	*waste = data

	fmt.Println("Data berhasil dimuat dari", filename)

	fmt.Println() // Spacing

	utils.PressToContinue()
	utils.ClearConsole()
}
