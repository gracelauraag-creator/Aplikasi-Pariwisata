# 🗺️ Aplikasi Pariwisata (Sadina & Grace)

Aplikasi Pariwisata adalah sebuah aplikasi berbasis teks (*Command-Line Interface*) yang dirancang untuk mengelola, mengurutkan, dan mencari data destinasi tempat wisata secara digital. Program ini dikembangkan menggunakan bahasa pemrograman **Go (Golang)** sebagai proyek Tugas Besar untuk mata kuliah **Algoritma Pemrograman** di Telkom University.

Aplikasi ini dibangun dengan arsitektur modular (subprogram) dan menerapkan batasan teknis ketat tanpa menggunakan kata kunci `break` atau `continue` pada perulangan biasa demi menghasilkan struktur logika algoritma yang bersih dan natural.

---

## 👥 Anggota Tim (IT-49-05)
1. **Sadina Dewi Amani S** - (NIM: 103032500092)
2. **Grace Laura Agnesia Nainggolan** - (NIM: 103032500128)

**Dosen Pengampu:** EEW  
**Program Studi:** S1 Teknologi Informasi, Telkom University  

---

## 🚀 Fitur Utama

Aplikasi ini membagi hak akses ke dalam 2 peran (*role*) utama:

### 👨‍💼 1. Menu Admin (Manajemen Data / CRUD)
* **Tambah Data Wisata**: Memasukkan destinasi baru ke dalam sistem (maksimal 100 data statis).
* **Tampilkan Semua Data**: Melihat seluruh daftar objek wisata yang telah terdaftar.
* **Edit Data Wisata**: Memperbarui informasi fasilitas, wahana, biaya, dan jarak tempuh tempat wisata.
* **Hapus Data Wisata**: Menghapus data tempat wisata dengan metode pergeseran (*shifting array*).

### 👤 2. Menu Pengguna (Akses Konten & Utilitas)
* **Lihat Daftar Tempat Wisata**: Menampilkan data tempat wisata yang tersedia.
* **Urutan Data (Sorting)**:
  * Berdasarkan **Biaya Tiket** (Ascending & Descending) — *Menggunakan Selection Sort*.
  * Berdasarkan **Jarak Lokasi** (Ascending & Descending) — *Menggunakan Insertion Sort*.
  * Berdasarkan **Nama Fasilitas** (Ascending & Descending) — *Menggunakan Selection Sort*.
* **Pencarian Data (Searching)**:
  * Cari berdasarkan **Nama Tempat** via *Sequential Search*.
  * Cari berdasarkan **Nama Tempat** via *Binary Search* (Aman menggunakan salinan array lokal tanpa merusak struktur data utama).
  * Cari berdasarkan **Fasilitas** atau **Wahana**.
* **🎟️ Fitur Bonus Rombongan**: Menghitung simulasi total tarif tiket berdasarkan jumlah anggota kelompok dengan skema **Diskon Otomatis 10%** jika rombongan berjumlah minimal 10 orang.

---

## 🛠️ Atribut Data (Struct Wisata)

Setiap tempat wisata dibungkus ke dalam sebuah tipe data bentukan (*struct*) dengan komponen:
* `nama` (string): Nama unik tempat wisata.
* `fasilitas` (string): Fasilitas penunjang utama.
* `wahana` (string): Wahana hiburan ikonik di lokasi.
* `biaya` (int): Harga tiket masuk per orang (Rp).
* `jarak` (int): Jarak tempuh menuju lokasi (Km).

---

## 💻 Cara Menjalankan Program

### Prasyarat
Pastikan komputer kamu sudah terinstal [Go compiler](https://go.dev/doc/install).

### Langkah-langkah
1. Klon repositori ini ke komputer lokal kamu:
   ```bash
   git clone [https://github.com/username-kamu/nama-repositori.git](https://github.com/username-kamu/nama-repositori.git)
