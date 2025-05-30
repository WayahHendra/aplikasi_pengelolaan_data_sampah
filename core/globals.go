package core

// Struktur data untuk pengelolaan sampah
type Waste struct {
	ID              string  // ID sampah
	WasteType       string  // Jenis sampah
	RecyclingMethod string  // Metode daur ulang
	Quantity        float64 // Jumlah sampah (dalam kg)
	Location        string  // Lokasi pengumpulan sampah
	Status          string  // Status daur ulang (sudah/belum)
	CreatedAt       string  // Waktu pembuatan data
	UpdatedAt       string  // Waktu pembaruan data
}

// Daftar variabel terkait data sampah
var (
	WasteData       []Waste            // Slice untuk menyimpan data sampah
	TriggerShowData bool    = true     // Trigger untuk menggunakan PressToContinue dan ClearConsole di fungsi ReadWaste
	Version         string  = "v2.4.4" // Versi aplikasi
	SwitchLanguage  bool    = false    // Trigger untuk menggunakan bahasa Inggris
)

// Prioritas port
var PRIORITY_PORTS = []int{5173, 5174, 5175, 5176, 5177}

// Konstanta api key  untuk middleware
const API_KEY = "1f0e7927-4e17-4abf-b457-5a691b4845f1"
