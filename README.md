# 🚀 Aplikasi Manajemen Startup Sederhana

## 📌 Deskripsi Singkat

Aplikasi ini adalah program berbasis **Go (Golang)** yang dirancang untuk membantu pengguna dalam mengelola data startup dan anggota timnya. Program ini bersifat **interaktif berbasis CLI (Command Line Interface)** dan memiliki fitur untuk memasukkan, menghapus, mengedit, mengurutkan, dan mencari data startup.

---

## 👨‍💻 Dibuat Oleh

- **Nama**: Sandres Sitorus
- **NIM**: 103012400100
- **Nama**: Alif Yaasin
- **NIM**: 103012400063
- **Versi**: Beta v0.0.7  

---

## 🎯 Tujuan Tugas

Tugas ini bertujuan untuk:
- Menerapkan **struktur data array statis** dalam bahasa Go
- Mengimplementasikan operasi **CRUD (Create, Read, Update, Delete)**
- Menerapkan **algoritma pencarian dan pengurutan**
- Melatih kemampuan pemrograman terstruktur dan modular
- Memberi pengalaman membuat program berbasis teks interaktif

---

## 🧩 Fitur Utama

### 🔧 Sebagai Admin:
- Tambah, ubah, dan hapus data startup
- Tambah dan hapus anggota tim startup
- Urutkan data berdasarkan tahun atau pendanaan
- Lihat laporan per bidang usaha
- Pencarian berdasarkan nama atau bidang

### 👤 Sebagai Pengguna:
- Tampilkan daftar startup
- Lihat anggota dari startup tertentu
- Akses fitur pencarian dan laporan

---

## 💾 Struktur Data yang Digunakan

```go
type startup struct {
    idStartup int
    nama      string
    pendanaan int
    tahun     int
    bidang    string
}

type anggota struct {
    idStartupAnggota int
    idAnggota        int
    namaAnggota      string
    peranAnggota     string
}
