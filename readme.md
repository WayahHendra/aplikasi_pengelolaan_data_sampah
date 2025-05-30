# Aplikasi Pengelolaan Data Sampah dan Daur Ulang

## Deskripsi
Aplikasi ini dirancang untuk mencatat dan mengelola data pengelolaan sampah serta proses daur ulang di suatu wilayah. Dengan fitur yang lengkap, aplikasi ini dapat digunakan oleh petugas pengelolaan sampah atau masyarakat yang ingin mencatat kontribusi mereka dalam daur ulang. Data utama yang dikelola meliputi jenis sampah, jumlah sampah yang dikumpulkan, dan metode daur ulang yang digunakan.

## Fitur Utama
1. **CRUD Data Sampah**:
   - Tambah, ubah, dan hapus data jenis sampah serta jumlah sampah yang dikumpulkan.
2. **Pencatatan Daur Ulang**:
   - Catat metode daur ulang yang digunakan untuk setiap jenis sampah.
3. **Pencarian Data**:
   - Cari data sampah berdasarkan jenis menggunakan Sequential dan Binary Search.
4. **Pengurutan Data**:
   - Urutkan data sampah berdasarkan jumlah atau jenis menggunakan Selection dan Insertion Sort.
5. **Statistik**:
   - Tampilkan statistik jumlah sampah yang dikumpulkan dan jumlah yang berhasil didaur ulang.
6. **Dukungan Bahasa**:
   - Aplikasi mendukung Bahasa Indonesia dan Bahasa Inggris.

## Spesifikasi Teknis
- **Bahasa Pemrograman**: Go
- **Format Data**: JSON
- **Mode Operasi**: CLI (Command Line Interface) dan GUI (Graphical User Interface - dalam pengembangan)

## Cara Instalasi
1. Clone repositori ini:
   ```bash
   git clone https://github.com/WayahHendra/aplikasi_pengelolaan_data_sampah.git
   ```
2. Masuk ke direktori proyek:
   ```bash
   cd aplikasi_pengelolaan_data_sampah/cmd
   ```
3. Jalankan aplikasi:
   ```bash
   go run main.go
   ```

## Cara Penggunaan
### Mode CLI
Setelah aplikasi dijalankan, pengguna dapat memilih menu berikut:
1. Tambah data sampah
2. Tampilkan data sampah
3. Ubah data sampah
4. Hapus data sampah
5. Cari data sampah
6. Urutkan data sampah
7. Tampilkan statistik
8. Simpan semua data
9. Muat semua data
10. Muat data berdasarkan tanggal
11. Ganti bahasa
12. Keluar dari aplikasi

### Mode GUI
Mode GUI sedang dalam tahap pengembangan dan akan tersedia di versi mendatang.

## Struktur Proyek
```
trash-app
├── cli/
│   ├── cli-logic/
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── update.go
│   │   ├── search.go
│   │   ├── sort.go
│   │   ├── statistics.go
│   │   ├── load.go
│   │   ├── save.go
│   │   └── switchlanguage.go
│   └── cli-menus/
│       ├── menu.go
│       └── selection.go
├── cmd/
│   └── main.go
├── core/
│   ├── crud.go
│   ├── globals.go
│   ├── shared.go
│   ├── storage.go
│   └── search.go
├── data/
│   └── dd-mm-yyyy.json
├── gui/
│   ├── gui.main.go
│   └── api/
│       ├── controller/
│       ├── module/
│       ├── router/
│       ├── service/
│       └── middleware/
│           ├── apikey.middleware.go
│           └── cors.middleware.go
├── utils/
│   ├── console.go
│   ├── logcolors.go
│   ├── port.go 
│   ├── stringutil.go
│   └── uuidv4.go
├── LICENSE
└── readme.md
```

## Kontribusi
Kami menerima kontribusi dari siapa saja! Jika Anda menemukan bug atau memiliki saran untuk fitur baru, silakan buka issue atau kirim pull request.

## Lisensi
Proyek ini dilisensikan di bawah MIT License. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.

## Kontak
Untuk pertanyaan lebih lanjut, silakan hubungi:
- **Email**: wayahhendrask@gmail.com || yogada017@gmail.com
- **GitHub**: [WayahHendra](https://github.com/WayahHendra) || [YogaDwiAlvian](https://github.com/AphidZ)
