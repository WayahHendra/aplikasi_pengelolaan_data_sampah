package cli_logic

import (
	"fmt"
	"trash-app/core"
	"trash-app/utils"
)

// CreateWaste memungkinkan pengguna untuk menambahkan data sampah baru.
func CreateWaste() {
	utils.ClearConsole()

	var (
		wasteType       string  // Jenis sampah
		recyclingMethod string  // Metode daur ulang
		quantity        float64 // Jumlah sampah (dalam kg)
		location        string  // Lokasi pengumpulan
		status          string  // Status daur ulang (sudah/belum)
	)

	var (
		count    int          // Jumlah data yang akan ditambahkan
		tempData []core.Waste // Slice sementara untuk menyimpan data baru
		waste    core.Waste   // Variabel sementara untuk data sampah
	)

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
		fmt.Println("How many waste data do you want to add?")
		fmt.Print("Enter the amount [positive integer]: ")
	} else {
		fmt.Println("Berapa banyak data sampah yang ingin ditambahkan?")
		fmt.Print("Masukkan jumlah [bilangan bulat positif]: ")
	}
	fmt.Scan(&count)
	if count == -1 {
		return // Kembali ke menu utama
	}

	// Validasi jumlah data
	if count <= 0 {
		utils.ClearConsole()

		if core.SwitchLanguage {
			fmt.Println("Invalid amount of data. Minimum 1.")
		} else {
			fmt.Println("Jumlah data tidak valid. Minimal 1.")
		}

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to try again...")
		} else {
			utils.PressToContinue("Tekan Enter untuk mencoba lagi...")
		}

		utils.ClearConsole()

		CreateWaste() // Panggil kembali fungsi CreateWaste untuk mencoba lagi

		return
	}

	utils.ClearConsole()

	// Loop untuk menambahkan data sebanyak jumlah yang dimasukkan
	for i := 0; i < count; i++ {
		core.ShowRecycleTypeTable() // Tampilkan tabel data daur ulang yang tersedia

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("=========================================")
			fmt.Println("[-1] Back to main menu [Cancel operation]")
			fmt.Println("================ Data", i+1, "=================")
		} else {
			fmt.Println("=============================================")
			fmt.Println("[-1] Kembali ke menu utama [Batalkan operasi]")
			fmt.Println("================= Data", i+1, "===================")
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

		// Tambahkan data ke slice sementara
		tempData = append(tempData, core.CreateData(wasteType, recyclingMethod, quantity, location, status))

		fmt.Println()

		if core.SwitchLanguage {
			fmt.Println("Data", i+1, "has been successfully added.")
		} else {
			fmt.Println("Data", i+1, "berhasil ditambahkan.")
		}

		if core.SwitchLanguage {
			utils.PressToContinue("Press Enter to continue...")
		} else {
			utils.PressToContinue("Tekan Enter untuk melanjutkan...")
		}

		utils.ClearConsole()
	}

	if core.SwitchLanguage {
		fmt.Println("Here is the data that has just been added:")
		fmt.Println("========================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Type", "Recycling Method", "Quantity", "Location", "Status")
		fmt.Println("||====================================================================================================||")
	} else {
		fmt.Println("Berikut adalah data yang baru ditambahkan:")
		fmt.Println("========================================================================================================")
		fmt.Printf("|| %-3s || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", "No", "Jenis", "Metode Daur Ulang", "Jumlah", "Lokasi", "Status")
		fmt.Println("||====================================================================================================||")
	}
	// Tampilkan data yang baru ditambahkan
	for i := 0; i < len(tempData); i++ {
		waste = tempData[i] // Ambil data dari slice sementara
		fmt.Printf("|| %-3d || %-15s || %-20s || %-12s || %-20s || %-8s ||\n", i+1, waste.WasteType, waste.RecyclingMethod, fmt.Sprint(waste.Quantity, " kg"), waste.Location, waste.Status)
	}
	fmt.Println("========================================================================================================")

	core.WasteData = append(core.WasteData, tempData...) // Tambahkan data baru ke dalam slice utama

	if core.SwitchLanguage {
		utils.PressToContinue("Press Enter to back to the main menu...")
	} else {
		utils.PressToContinue("Tekan Enter untuk kembali ke menu utama...")
	}

	utils.ClearConsole()
}
