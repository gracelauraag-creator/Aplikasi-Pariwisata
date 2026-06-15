package main

import "fmt"

const MAX = 100

type Wisata struct {
	nama      string
	fasilitas string
	wahana    string
	biaya     int
	jarak     int
}

var data [MAX]Wisata
var n int


func tambahData(arr *[MAX]Wisata, total *int) {
	if *total < MAX {
		var sisa string
		fmt.Scanln(&sisa)

		fmt.Print("Nama wisata : ")
		fmt.Scan(&arr[*total].nama)

		fmt.Print("Fasilitas   : ")
		fmt.Scan(&arr[*total].fasilitas)

		fmt.Print("Wahana      : ")
		fmt.Scan(&arr[*total].wahana)

		fmt.Print("Biaya       : ")
		fmt.Scan(&arr[*total].biaya)

		fmt.Print("Jarak       : ")
		fmt.Scan(&arr[*total].jarak)

		*total++
		fmt.Println("Data berhasil ditambahkan")
	} else {
		fmt.Println("Kapasitas data sudah penuh!")
	}
}

func tampilData(arr [MAX]Wisata, total int) {
	if total == 0 {
		fmt.Println("Belum ada data tempat wisata.")
	} else {
		for i := 0; i < total; i++ {
			fmt.Println("---------------------------------")
			fmt.Println("Nama      :", arr[i].nama)
			fmt.Println("Fasilitas :", arr[i].fasilitas)
			fmt.Println("Wahana    :", arr[i].wahana)
			fmt.Println("Biaya     :", arr[i].biaya)
			fmt.Println("Jarak     :", arr[i].jarak)
		}
		fmt.Println("---------------------------------")
	}
}

func editData(arr *[MAX]Wisata, total int) {
	var nama string
	var metode, idx int

	fmt.Print("Nama tempat wisata yang akan diedit : ")
	fmt.Scan(&nama)

	fmt.Println("Pilih metode pencarian untuk edit data:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Urutan alfabetis nama)")
	fmt.Print("Pilih (1-2): ")
	fmt.Scan(&metode)

	if metode == 1 {
		idx = sequentialSearchNama(*arr, total, nama)
	} else {
		idx = binarySearchNama(arr, total, nama)
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Print("Fasilitas baru : ")
		fmt.Scan(&arr[idx].fasilitas)

		fmt.Print("Wahana baru    : ")
		fmt.Scan(&arr[idx].wahana)

		fmt.Print("Biaya baru     : ")
		fmt.Scan(&arr[idx].biaya)

		fmt.Print("Jarak baru     : ")
		fmt.Scan(&arr[idx].jarak)

		fmt.Println("Data berhasil diubah")
	}
}

func hapusData(arr *[MAX]Wisata, total *int) {
	var nama string
	var metode, idx int

	fmt.Print("Nama tempat wisata yang akan dihapus : ")
	fmt.Scan(&nama)

	fmt.Println("Pilih metode pencarian untuk hapus data:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Urutan alfabetis nama)")
	fmt.Print("Pilih (1-2): ")
	fmt.Scan(&metode)

	if metode == 1 {
		idx = sequentialSearchNama(*arr, *total, nama)
	} else {
		idx = binarySearchNama(arr, *total, nama)
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		for i := idx; i < *total-1; i++ {
			arr[i] = arr[i+1]
		}
		*total--
		fmt.Println("Data berhasil dihapus")
	}
}

// ==================== ALGORITMA SEARCHING ====================

func sequentialSearchNama(arr [MAX]Wisata, total int, nama string) int {
	var hasil int
	hasil = -1

	for i := 0; i < total; i++ {
		if arr[i].nama == nama {
			hasil = i
		}
	}
	return hasil
}

func sequentialSearchFasilitas(arr [MAX]Wisata, total int, fasilitas string) {
	var found bool
	found = false
	for i := 0; i < total; i++ {
		if arr[i].fasilitas == fasilitas {
			fmt.Println("- ", arr[i].nama)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearchWahana(arr [MAX]Wisata, total int, wahana string) {
	var found bool
	found = false
	for i := 0; i < total; i++ {
		if arr[i].wahana == wahana {
			fmt.Println("- ", arr[i].nama)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func urutNama(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var min int
	for i := 0; i < total-1; i++ {
		min = i
		for j := i + 1; j < total; j++ {
			if arr[j].nama < arr[min].nama {
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func binarySearchNama(arr *[MAX]Wisata, total int, nama string) int {
	var arrLokal [MAX]Wisata
	for i := 0; i < total; i++ {
		arrLokal[i] = arr[i]
	}
	urutNama(&arrLokal, total) 

	var kiri, kanan, tengah, hasil int
	kiri = 0
	kanan = total - 1
	hasil = -1

	for kiri <= kanan && hasil == -1 {
		tengah = (kiri + kanan) / 2
		if arrLokal[tengah].nama == nama 
			hasil = sequentialSearchNama(*arr, total, arrLokal[tengah].nama)
		} else if nama < arrLokal[tengah].nama {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return hasil
}

// ==================== ALGORITMA SORTING ====================

func selectionSortBiayaAsc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var min int
	for i := 0; i < total-1; i++ {
		min = i
		for j := i + 1; j < total; j++ {
			if arr[j].biaya < arr[min].biaya {
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func selectionSortBiayaDesc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var max int
	for i := 0; i < total-1; i++ {
		max = i
		for j := i + 1; j < total; j++ {
			if arr[j].biaya > arr[max].biaya {
				max = j
			}
		}
		temp = arr[i]
		arr[i] = arr[max]
		arr[max] = temp
	}
}

func insertionSortJarakAsc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var j int
	for i := 1; i < total; i++ {
		temp = arr[i]
		j = i - 1
		for j >= 0 && arr[j].jarak > temp.jarak {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func insertionSortJarakDesc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var j int
	for i := 1; i < total; i++ {
		temp = arr[i]
		j = i - 1
		for j >= 0 && arr[j].jarak < temp.jarak {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func selectionSortFasilitasAsc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var min int
	for i := 0; i < total-1; i++ {
		min = i
		for j := i + 1; j < total; j++ {
			if arr[j].fasilitas < arr[min].fasilitas {
				min = j
			}
		}
		temp = arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func selectionSortFasilitasDesc(arr *[MAX]Wisata, total int) {
	var temp Wisata
	var max int
	for i := 0; i < total-1; i++ {
		max = i
		for j := i + 1; j < total; j++ {
			if arr[j].fasilitas > arr[max].fasilitas {
				max = j
			}
		}
		temp = arr[i]
		arr[i] = arr[max]
		arr[max] = temp
	}
}

// ==================== 🎟️ FITUR TAMBAHAN KREATIF 🎟️ ====================

func hitungBiayaRombongan(arr [MAX]Wisata, total int) {
	var namaWisata string
	var jumlahOrang int

	if total == 0 {
		fmt.Println("Belum ada data tempat wisata terdaftar.")
		return
	}

	fmt.Print("Masukkan Nama Wisata tujuan rombongan: ")
	fmt.Scan(&namaWisata)

	idx := sequentialSearchNama(arr, total, namaWisata)

	if idx == -1 {
		fmt.Println("Tempat wisata tersebut tidak ditemukan.")
	} else {
		fmt.Print("Masukkan jumlah anggota rombongan (orang): ")
		fmt.Scan(&jumlahOrang)

		totalBiaya := arr[idx].biaya * jumlahOrang
		diskon := 0

		if jumlahOrang >= 10 {
			diskon = totalBiaya * 10 / 100 
			totalBiaya = totalBiaya - diskon
			fmt.Println("🎉 Selamat! Anda mendapatkan Potongan Diskon Rombongan sebesar 10%!")
		}

		fmt.Println("\n=========================================")
		fmt.Println("        RINCIAN BIAYA ROMBONGAN          ")
		fmt.Println("=========================================")
		fmt.Println("Destinasi Tujuan :", arr[idx].nama)
		fmt.Println("Harga Tiket/Orang: Rp", arr[idx].biaya)
		fmt.Println("Jumlah Anggota   :", jumlahOrang, "orang")
		fmt.Println("Potongan Diskon  : Rp", diskon)
		fmt.Println("-----------------------------------------")
		fmt.Println("TOTAL BAYAR      : Rp", totalBiaya)
		fmt.Println("=========================================")
	}
}

func main() {
	var role, pilih, jenis int
	var nama, fasilitas, wahana string
	var idx int

	role = 0
	for role != 3 {
		fmt.Println("\n=================================")
		fmt.Println("   APLIKASI PARIWISATA (SADINA & GRACE)   ")
		fmt.Println("=================================")
		fmt.Println("1. Masuk sebagai Admin")
		fmt.Println("2. Masuk sebagai Pengguna")
		fmt.Println("3. Keluar dari Aplikasi")
		fmt.Print("Pilih Peran Anda (1-3): ")
		fmt.Scan(&role)

		if role == 1 {
			pilih = 0
			for pilih != 5 {
				fmt.Println("\n--- MENU ADMIN ---")
				fmt.Println("1. Tambah Data Wisata")
				fmt.Println("2. Tampilkan Semua Data")
				fmt.Println("3. Edit Data Wisata")
				fmt.Println("4. Hapus Data Wisata")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilih Menu Admin: ")
				fmt.Scan(&pilih)

				if pilih == 1 {
					tambahData(&data, &n)
				} else if pilih == 2 {
					tampilData(data, n)
				} else if pilih == 3 {
					editData(&data, n)
				} else if pilih == 4 {
					hapusData(&data, &n)
				}
			}

		} else if role == 2 {
			pilih = 0
			for pilih != 9 {
				fmt.Println("\n--- MENU PENGGUNA ---")
				fmt.Println("1. Lihat Daftar Tempat Wisata")
				fmt.Println("2. Urutkan berdasarkan Biaya")
				fmt.Println("3. Urutkan berdasarkan Jarak")
				fmt.Println("4. Urutkan berdasarkan Fasilitas")
				fmt.Println("5. Cari berdasarkan Nama (Sequential)")
				fmt.Println("6. Cari berdasarkan Nama (Binary)")
				fmt.Println("7. Cari Tempat berdasarkan Fasilitas/Wahana")
				fmt.Println("8. 🎟️ Hitung Biaya Rombongan (Fitur Tambahan)")
				fmt.Println("9. Kembali ke Menu Utama")
				fmt.Print("Pilih Menu Pengguna: ")
				fmt.Scan(&pilih)

				if pilih == 1 {
					tampilData(data, n)
				} else if pilih == 2 {
					fmt.Println("1. Terkecil ke Terbesar (Ascending)")
					fmt.Println("2. Terbesar ke Terkecil (Descending)")
					fmt.Print("Pilih jenis urutan: ")
					fmt.Scan(&jenis)
					if jenis == 1 {
						selectionSortBiayaAsc(&data, n)
					} else {
						selectionSortBiayaDesc(&data, n)
					}
					fmt.Println("Data berhasil diurutkan!")
					tampilData(data, n)

				} else if pilih == 3 {
					fmt.Println("1. Terdekat ke Terjauh (Ascending)")
					fmt.Println("2. Terjauh ke Terdekat (Descending)")
					fmt.Print("Pilih jenis urutan: ")
					fmt.Scan(&jenis)
					if jenis == 1 {
						insertionSortJarakAsc(&data, n)
					} else {
						insertionSortJarakDesc(&data, n)
					}
					fmt.Println("Data berhasil diurutkan!")
					tampilData(data, n)

				} else if pilih == 4 {
					fmt.Println("1. A-Z (Ascending)")
					fmt.Println("2. Z-A (Descending)")
					fmt.Print("Pilih jenis urutan: ")
					fmt.Scan(&jenis)
					if jenis == 1 {
						selectionSortFasilitasAsc(&data, n)
					} else {
						selectionSortFasilitasDesc(&data, n)
					}
					fmt.Println("Data berhasil diurutkan!")
					tampilData(data, n)

				} else if pilih == 5 {
					fmt.Print("Masukkan nama wisata yang dicari: ")
					fmt.Scan(&nama)
					idx = sequentialSearchNama(data, n, nama)
					if idx == -1 {
						fmt.Println("Data tidak ditemukan")
					} else {
						fmt.Println("\n--- Hasil Pencarian ---")
						fmt.Println("Nama      :", data[idx].nama)
						fmt.Println("Fasilitas :", data[idx].fasilitas)
						fmt.Println("Wahana    :", data[idx].wahana)
						fmt.Println("Biaya     :", data[idx].biaya)
						fmt.Println("Jarak     :", data[idx].jarak)
					}

				} else if pilih == 6 {
					fmt.Print("Masukkan nama wisata yang dicari: ")
					fmt.Scan(&nama)
					idx = binarySearchNama(&data, n, nama)
					if idx == -1 {
						fmt.Println("Data tidak ditemukan")
					} else {
						fmt.Println("\n--- Hasil Pencarian (Binary Search) ---")
						fmt.Println("Nama      :", data[idx].nama)
						fmt.Println("Fasilitas :", data[idx].fasilitas)
						fmt.Println("Wahana    :", data[idx].wahana)
						fmt.Println("Biaya     :", data[idx].biaya)
						fmt.Println("Jarak     :", data[idx].jarak)
					}

				} else if pilih == 7 {
					fmt.Println("1. Cari lewat Fasilitas")
					fmt.Println("2. Cari lewat Wahana")
					fmt.Print("Pilih sub-kategori: ")
					fmt.Scan(&jenis)
					if jenis == 1 {
						fmt.Print("Masukkan nama fasilitas: ")
						fmt.Scan(&fasilitas)
						fmt.Println("Tempat wisata dengan fasilitas tersebut:")
						sequentialSearchFasilitas(data, n, fasilitas)
					} else {
						fmt.Print("Masukkan nama wahana: ")
						fmt.Scan(&wahana)
						fmt.Println("Tempat wisata dengan wahana tersebut:")
						sequentialSearchWahana(data, n, wahana)
					}
				} else if pilih == 8 {
					hitungBiayaRombongan(data, n)
				}
			}
		}
	}
	fmt.Println("Terima kasih telah menggunakan Aplikasi Pariwisata!")
}