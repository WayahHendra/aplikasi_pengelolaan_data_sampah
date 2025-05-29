package cli_logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"trash-app/core"
	"trash-app/utils"
)

// LoadWaste memungkinkan pengguna untuk memuat data sampah dari file JSON.
func LoadWaste(waste *[]core.Waste) {
	utils.ClearConsole()

	var filename string // Nama file yang akan dimuat

	filename = "../data/" + time.Now().Format("02-01-2006") + ".json"
	filename = utils.EnsureJSONExtension(filename) // Pastikan nama file memiliki ekstensi .json

	// Memuat data dari file
	data, err := core.LoadWasteFromFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		// Jika berhasil, data dimasukkan ke dalam slice
		*waste = data
	}

	utils.ClearConsole()
}

// LoadWasteByDate memungkinkan pengguna untuk memuat data sampah dari file JSON berdasarkan tanggal.
func LoadWasteByDate(waste *[]core.Waste) (inputDate string) {
    utils.ClearConsole()

    var filename, date string

	if core.SwitchLanguage {
		fmt.Println("Enter the date to load data, on that date")
		fmt.Println("Example date format: 02-01-2006")
	} else {
		fmt.Println("Masukkan tanggal untuk load data, pada tanggal tersebut")
		fmt.Println("Contoh format tanggal: 02-01-2006")
	}

    fmt.Printf("> ")
    fmt.Scan(&date)

    // Persiapkan nama file berdasarkan input
    filename = "../data/" + date + ".json"
    filename = utils.EnsureJSONExtension(filename)

    data, err := core.LoadWasteFromFile(filename)
    if err != nil {
        // Jika error, fallback ke data hari ini
        today := time.Now().Format("02-01-2006")
        fallbackFile := "../data/" + today + ".json"
        fallbackFile = utils.EnsureJSONExtension(fallbackFile)

        data, err = core.LoadWasteFromFile(fallbackFile)
        if err != nil {
            // Kalau data fallback juga tidak ada, tampilkan error
            if core.SwitchLanguage {
                fmt.Println("No data available for the provided date and today.")
            } else {
                fmt.Println("Data tidak tersedia untuk tanggal yang dimasukkan dan hari ini.")
            }
            return "" // atau bisa juga return today
        }

        *waste = data
        inputDate = today

        fmt.Println()
        if core.SwitchLanguage {
            fmt.Println("Data not found for the given date. Loading today's data instead...")
            time.Sleep(2 * time.Second) // Menunggu 2 detik
            fmt.Println("Data successfully loaded from", fallbackFile)
        } else {
            fmt.Println("Data tidak ditemukan untuk tanggal yang dimasukkan. Memuat data hari ini...")
            time.Sleep(2 * time.Second) // Menunggu 2 detik
            fmt.Println("Data berhasil dimuat dari", fallbackFile)
        }
    } else {
        *waste = data
        inputDate = date

        fmt.Println()
        if core.SwitchLanguage {
            fmt.Println("Data found for the given date.")
            time.Sleep(2 * time.Second)
            fmt.Println("Data successfully loaded from", filename)
        } else {
            fmt.Println("Data ditemukan untuk tanggal yang dimasukkan.")
            time.Sleep(2 * time.Second)
            fmt.Println("Data berhasil dimuat dari", filename)
        }
    }

    if core.SwitchLanguage {
        utils.PressToContinue("Press Enter to continue...")
    } else {
        utils.PressToContinue("Tekan Enter untuk melanjutkan...")
    }

    utils.ClearConsole()
    return inputDate
}

// LoadAllWaste memuat semua data sampah dari semua file JSON di direktori data.	
func LoadAllWaste(waste *[]core.Waste) {
    utils.ClearConsole()

    var dataFolder = "../data/"
    var allWaste []core.Waste

    files, err := os.ReadDir(dataFolder)
    if err != nil {
        if core.SwitchLanguage {
            fmt.Printf("Failed to read data directory: %v\n", err)
        } else {
            fmt.Printf("Gagal membaca direktori data: %v\n", err)
        }
        return
    }

    for i := 0; i < len(files); i++ {
        file := files[i]

        if file.IsDir() {
            continue
        }

        fileName := file.Name()
        if !strings.HasSuffix(fileName, ".json") {
            continue
        }

        filePath := filepath.Join(dataFolder, fileName)
        fileContent, err := os.ReadFile(filePath)
        if err != nil {
            if core.SwitchLanguage {
                fmt.Printf("Failed to read file %s: %v\n", filePath, err)
            } else {
                fmt.Printf("Gagal membaca file %s: %v\n", filePath, err)
            }
            continue
        }

        var wasteData []core.Waste
        if err := json.Unmarshal(fileContent, &wasteData); err != nil {
            if core.SwitchLanguage {
                fmt.Printf("Failed to unmarshal file %s: %v\n", filePath, err)
            } else {
                fmt.Printf("Gagal memuat file %s: %v\n", filePath, err)
            }
            continue
        }

        allWaste = append(allWaste, wasteData...)
    }

    *waste = allWaste

    if core.SwitchLanguage {
        fmt.Println("Trying to load all waste data...")
        time.Sleep(2 * time.Second) // Menunggu 2 detik
        fmt.Println("All waste data loaded successfully.")
    } else {
        fmt.Println("Mencoba memuat semua data sampah...")
        time.Sleep(2 * time.Second) // Menunggu 2 detik
        fmt.Println("Semua data sampah berhasil dimuat.")
    }

    if core.SwitchLanguage {
        utils.PressToContinue("Press Enter to continue...")
    } else {
        utils.PressToContinue("Tekan Enter untuk melanjutkan...")
    }

    utils.ClearConsole()
}