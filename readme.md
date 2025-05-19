# Aplikasi Daur Ulang Sampah

## Overview
Trash App adalah aplikasi berbasis CLI dan GUI yang dirancang untuk membantu pengelolaan data sampah, termasuk menambah, membaca, mengubah, menghapus, mencari, mengurutkan, mengganti bahasa, dan menyimpan data sampah. Aplikasi ini bertujuan untuk meningkatkan kesadaran akan pengelolaan sampah yang lebih baik.

## Features
- **GUI Placeholder**: GUI belum diimplementasikan, tetapi tersedia sebagai placeholder.
- **CLI Menu**: Menyediakan berbagai fitur seperti CRUD, pencarian, pengurutan, dan lainnya.
- **Pengelolaan Data Sampah**: Mendukung berbagai jenis sampah dan metode daur ulang.
- **Pemilihan Bahasa**: Pengguna dapat mengganti ke bahasa Inggris maupun bahasa Indonesia.
- **Penyimpanan dan Pemulihan Data**: Data dapat disimpan ke file JSON dan dimuat kembali.

## Installation
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
   go run .
   ```

## Usage
Saat menjalankan aplikasi, Anda akan diminta untuk memilih mode:
- **GUI Mode**: Tekan `y` (placeholder, belum diimplementasikan).
- **CLI Mode**: Tekan `n` untuk menggunakan mode CLI.

### GUI Menu
GUI belum diimplementasikan!

### CLI Menu
Berikut adalah daftar menu yang tersedia di mode CLI:
1. Tambah data
2. Tampilkan semua data
3. Ubah data
4. Hapus data
5. Cari data (Not Implemented yet!)
6. Urutkan data
7. Catat proses daur ulang (Not Implemented yet!)
8. Tampilkan statistik (Not Implemented yet!)
9. Simpan data
10. Muat data
11. Ganti bahasa
12. Keluar program

## Project Structure
```
trash-app
├── cli/
│   ├── cli-logic/
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── load.go
│   │   ├── record.go
│   │   ├── read.go
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
│   ├── gui-logic/
│   │   ├── create.go
│   │   ├── delete.go
│   │   ├── read.go
│   │   └── update.go
│   └── gui-menus/
│       └── menu.go
├── utils/
│   ├── console.go
│   ├── input.go
│   └── stringutil.go
└── go.mod
```

## Contributing
Kontribusi sangat diterima! Jika Anda memiliki saran atau ingin melaporkan bug, silakan buka issue atau kirim pull request.

## License
Proyek ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.
