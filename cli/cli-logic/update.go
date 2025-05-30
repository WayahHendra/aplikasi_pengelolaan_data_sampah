package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// UpdateWaste memungkinkan pengguna untuk mengubah data sampah yang sudah ada.
func UpdateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string  // Jenis sampah
		recyclingMethod string  // Metode daur ulang
		quantity        float64 // Jumlah sampah (dalam kg)
		location        string  // Lokasi pengumpulan
		status          string  // Status daur ulang (sudah/belum)
	)

	var (
		confirmUpdate string // Konfirmasi pengubahan data
		index         int    // Indeks data yang akan diubah
	)

	core.TriggerShowData = false // Nonaktifkan trigger untuk membersihkan konsol
	ReadWaste("")                // Menampilkan tabel data
	core.TriggerShowData = true  // Aktifkan kembali trigger

	// Validasi apakah ada data yang tersedia
	if len(core.WasteData) == 0 {
		return // Jika tidak ada data, keluar dari fungsi
	}

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Println("=========================================")
		fmt.Println("[-1] Back to main menu [Cancel operation]")
		fmt.Println("=========================================")
	} else {
		fmt.Println("=============================================")
		fmt.Println("[-1] Kembali ke menu utama [Batalkan operasi]")
		fmt.Println("=============================================")
	}

	fmt.Println()

	if core.SwitchLanguage {
		fmt.Print("Enter the number of the data you want to update: ")
	} else {
		fmt.Print("Masukkan nomor data yang ingin diubah: ")
	}
	fmt.Scan(&index)
	if index == -1 {
		return // Kembali ke menu utama
	}

	index-- // Ubah ke indeks berbasis 0

	fmt.Println()

	// Validasi indeks
	if index < 0 || index >= len(core.WasteData) {
		utils.ClearConsole()

		if core.SwitchLanguage {
			fmt.Println("Invalid data number. Try again.")
		} else {
			fmt.Println("Nomor data tidak valid. Coba lagi.")
		}

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to continue...")
		} else {
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}

		UpdateWaste() // Panggil kembali fungsi UpdateWaste untuk mencoba lagi
	}

	utils.ClearConsole()

	if core.SwitchLanguage {
		fmt.Print("Are you sure you want to update the data? [1] Yes, [2] No: ")
	} else {
		fmt.Print("Yakin ingin mengubah data? [1] Ya, [2] Tidak: ")
	}
	fmt.Scan(&confirmUpdate)
	if confirmUpdate == "-1" {
		return // Kembali ke menu utama
	}

	if confirmUpdate == "1" {
		utils.ClearConsole()

		core.ShowRecycleTypeTable() // Menampilkan jenis-jenis sampah daur ulang

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Data to be changed:")
			fmt.Println("========================================================================================================")
			fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Type", "Recycling Method", "Quantity", "Location", "Status")
			fmt.Println("||====================================================================================================||")
		} else {
			fmt.Println("Data yang akan diubah:")
			fmt.Println("========================================================================================================")
			fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
			fmt.Println("||====================================================================================================||")
		}
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", index+1, core.WasteData[index].WasteType, core.WasteData[index].RecyclingMethod, fmt.Sprint(core.WasteData[index].Quantity, " kg"), core.WasteData[index].Location, core.WasteData[index].Status)
		fmt.Println("========================================================================================================")

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("=========================================")
			fmt.Println("[-1] Back to main menu [Cancel operation]")
			fmt.Println("=========================================")
		} else {
			fmt.Println("=============================================")
			fmt.Println("[-1] Kembali ke menu utama [Batalkan operasi]")
			fmt.Println("=============================================")
		}

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Print("Enter waste type: ")
		} else {
			fmt.Print("Masukkan jenis sampah: ")
		}
		fmt.Scan(&wasteType)
		if wasteType == "-1" {
			return // Kembali ke menu utama
		}

		if core.SwitchLanguage {
			fmt.Print("Enter recycling method: ")
		} else {
			fmt.Print("Masukkan metode daur ulang: ")
		}
		fmt.Scan(&recyclingMethod)
		if recyclingMethod == "-1" {
			return // Kembali ke menu utama
		}

		if core.SwitchLanguage {
			fmt.Print("Enter waste quantity [kg]: ")
		} else {
			fmt.Print("Masukkan jumlah sampah [kg]: ")
		}
		fmt.Scan(&quantity)
		if quantity == -1 {
			return // Kembali ke menu utama
		}

		if core.SwitchLanguage {
			fmt.Print("Enter collection location: ")
		} else {
			fmt.Print("Masukkan lokasi pengumpulan: ")
		}
		fmt.Scan(&location)
		if location == "-1" {
			return // Kembali ke menu utama
		}

		if core.SwitchLanguage {
			fmt.Print("Enter recicling status. [1] Complete, [2] Incomplete: ")
		} else {
			fmt.Print("Masukkan status daur ulang. [1] Sudah, [2] Belum: ")
		}
		fmt.Scan(&status)
		if status == "-1" {
			return // Kembali ke menu utama
		}

		// Gunakan core.CreateData untuk normalisasi dan validasi data baru
		core.WasteData[index] = core.CreateData(wasteType, recyclingMethod, quantity, location, status)

		if core.SwitchLanguage {
			fmt.Println("Data successfully updated.")
		} else {
			fmt.Println("Data berhasil diubah.")
		}
	} else if confirmUpdate == "2" {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Failed to update data.")
		} else {
			fmt.Println("Gagal mengubah data.")
		}
	} else {
		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Invalid input!")
		} else {
			fmt.Println("Input tidak valid!")
		}
	}

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to return to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk melanjutkan...")
	}

	utils.ClearConsole()
}
