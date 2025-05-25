package core

// Struktur data untuk pengelolaan sampah
type Waste struct {
	WasteType       string  // Jenis sampah
	RecyclingMethod string  // Metode daur ulang
	Quantity        float64 // Jumlah sampah (dalam kg)
	Location        string  // Lokasi pengumpulan sampah
	Status          string  // Status daur ulang (sudah/belum)
}

// Daftar variabel terkait data sampah
var (
	WasteData       []Waste            // Slice untuk menyimpan data sampah
	TriggerShowData bool    = true     // Trigger untuk menggunakan PressToContinue dan ClearConsole di fungsi ReadWaste
	Version         string  = "v1.2.5" // Versi aplikasi
	SwitchLanguage  bool    = false    // Trigger untuk menggunakan bahasa Inggris
)
