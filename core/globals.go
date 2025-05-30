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
	Version         string  = "v2.3.3" // Versi aplikasi
	SwitchLanguage  bool    = false    // Trigger untuk menggunakan bahasa Inggris
)
