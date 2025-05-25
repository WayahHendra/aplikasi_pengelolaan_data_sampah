# Aplikasi Daur Ulang Sampah

## Overview
Trash App adalah aplikasi berbasis CLI dan GUI yang dirancang untuk membantu pengelolaan data sampah, termasuk menambah, membaca, mengubah, menghapus, mencari, mengurutkan, mengganti bahasa, dan menyimpan data sampah. Aplikasi ini bertujuan untuk meningkatkan kesadaran akan pengelolaan sampah yang lebih baik.

## Features
- **GUI Placeholder**: GUI belum diimplementasikan, sedang dalam tahap pengembangan!.
- **CLI Menu**: Menyediakan berbagai fitur seperti CRUD, pencarian, pengurutan, dan lainnya.
- **Pengelolaan Data Sampah**: Mendukung berbagai jenis sampah dan metode daur ulang.
- **Pemilihan Bahasa**: Pengguna dapat mengganti ke bahasa Inggris maupun bahasa Indonesia.
- **Penyimpanan dan Pemulihan Data**: Data dapat disimpan ke file JSON dan dimuat kembali.

## Installation
1. Clone repositori ini:
   ```bash
   git clone https://github.com/WayahHendra/aplikasi_pengelolaan_data_sampah.git
   ```

## Running
1. Masuk ke direktori proyek:
   ```bash
   cd aplikasi_pengelolaan_data_sampah/cmd
   ```

2. Jalankan aplikasi dengan salah satu perintah berikut:
   ```bash
   go run main.go
   ```
   atau
   ```bash
   go run .
   ```

3. Ikuti petunjuk di terminal untuk memilih mode dan menggunakan aplikasi.

## Usage
Saat menjalankan aplikasi, Anda akan diminta untuk memilih mode:
- **GUI Mode**: Tekan `y` (placeholder, belum diimplementasikan (development)).
- **CLI Mode**: Tekan `n` untuk menggunakan mode CLI.

### GUI Menu
GUI dalam tahap pengembangan!

### CLI Menu
Berikut adalah daftar menu yang tersedia di mode CLI:
[1] Tambah data
[2] Tampilkan semua data
[3] Ubah data
[4] Hapus data
[5] Cari data -> Not Implemented yet!
[6] Urutkan data
[7] Catat proses daur ulang -> Not Implemented yet!
[8] Tampilkan statistik -> Not Implemented yet!
[9] Simpan data -> Bug
[10] Muat data -> Bug
[11] Ganti bahasa
[-1] Keluar program

## Project Structure
```
trash-app
├── cli/
│   ├── cli-logic/
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── exitprogram.go
│   │   ├── load.go
│   │   ├── read.go
│   │   ├── record.go
│   │   ├── save.go
│   │   ├── search.go
│   │   ├── sort.go
│   │   ├── statistics.go
│   │   ├── switchlanguage.go
│   │   └── update.go
│   └── cli-menus/
│       ├── menu.go
│       └── selection.go
├── cmd/
│   └── main.go
├── core/
│   ├── crud.go
│   ├── globals.go
│   ├── search.go
│   ├── shared.go
│   ├── sort.go
│   └── storage.go
├── gui/
│   ├── api/
│   │   ├── controller/
│   │   │   └── waste.controllers.go
│   │   ├── module/
│   │   │   └── waste.modules.go
│   │   ├── router/
│   │   │   └── waste.routers.go
│   │   └── service/
│   │       └── waste.services.go
│   └── gui.main.go  
├── utils/
│   ├── console.go
│   └── stringutil.go
└── go.mod
```

## Contributing
Kontribusi sangat diterima! Jika Anda memiliki saran atau ingin melaporkan bug, silakan buka issue atau kirim pull request.

## License
Proyek ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.
