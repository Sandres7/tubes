package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const NMAX int = 1000

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

type tabelAnggota [NMAX]anggota
type tabelStartup [NMAX]startup

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func countdownDetik(detik int) {
	for i := detik; i > 0; i-- {
		time.Sleep(1 * time.Second)
	}
}

func menuAwal() {
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Printf("║%3sSelamat Datang Di Aplikasi Manajemen Startup Sederhana%-5s║\n", "", "")
	fmt.Printf("║%-24s %-37s║\n", "", "beta.0.0.7")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ 1. Masukkan data Startup Baru                                ║\n")
	fmt.Printf("║ 2. Load data startup yang sudah ada                          ║\n")
	fmt.Printf("║ 3. Change Log                                                ║\n")
	fmt.Printf("║ 4. Exit                                                      ║\n")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
}

func tampilkanMenuUtama() {
	clear()
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║    🚀 APLIKASI MANAJEMEN STARTUP 🚀  ║")
	fmt.Println("║        Dibuat oleh: Sandres Sitorus  ║")
	fmt.Println("║                     Alif Yasin       ║")
	fmt.Println("║            103012400100              ║")
	fmt.Println("║            103012400063              ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. 📁 Kelola Data Startup             ║")
	fmt.Println("║2. 👥 Kelola Tim Startup              ║")
	fmt.Println("║3. 🔍 Cari Startup                    ║")
	fmt.Println("║4. 📊 Urutkan Daftar Startup          ║")
	fmt.Println("║5. 🧾 Laporan Bidang Usaha            ║")
	fmt.Println("║6. ❌ Keluar Aplikasi                 ║")
	fmt.Println("║9. ⚙️ Pilih Role                      ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func menuPengguna() {
	clear()
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║    🚀 APLIKASI MANAJEMEN STARTUP 🚀  ║")
	fmt.Println("║        Dibuat oleh: Sandres Sitorus  ║")
	fmt.Println("║                     Alif Yasin       ║")
	fmt.Println("║            103012400100              ║")
	fmt.Println("║            103012400063              ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. 📁 Tampilkan Startup               ║")
	fmt.Println("║2. 👥 Tampilkan Anggota Startup       ║")
	fmt.Println("║3. 🔍 Cari Startup                    ║")
	fmt.Println("║4. 📊 Urutkan Daftar Startup          ║")
	fmt.Println("║5. 🧾 Laporan Bidang Usaha            ║")
	fmt.Println("║6. ❌ Keluar Aplikasi                 ║")
	fmt.Println("║9. ⚙️ Pilih Role                      ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func menuKelolaDataStartup(tabel *tabelStartup, tabAnggota *tabelAnggota) {
	clear()
	var pilih int
	var lokasi int
	var nama string
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║      🚀 KELOLA DATA STARTUP 🚀       ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. ➕ Tambah Data Startup             ║")
	fmt.Println("║2. 📝 Ubah Data Startup               ║")
	fmt.Println("║3. 🗑️ Hapus Data Startup              ║")
	fmt.Println("║4. 📃 Lihat Semua Startup             ║")
	fmt.Println("║5. 🔙 Kembali ke Menu Utama           ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		sortIndexStartup(tabel)
		printData(*tabel)
		tambahStartup(tabel)
		fmt.Println("Data berhasil ditambah !")
		fmt.Println("Akan Kembali ke Menu Utama")
		countdownDetik(2)
	case 2:
		sortIndexStartup(tabel)
		printData(*tabel)
		fmt.Println("Masukkan Nama Startup yang ingin diubah (Masukkan - untuk cancel)")
		fmt.Scan(&nama)
		if nama == "-" {
			return
		}
		nama = strings.ToLower(nama)
		lokasi = sequentialSearchStartup(*tabel, nama)
		if lokasi == -1 {
			fmt.Println("Data tidak ditemukan")
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(2)
		} else {
			editStartup(tabel, lokasi)
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(2)
		}
	case 3:
		sortIndexStartup(tabel)
		printData(*tabel)
		fmt.Println("Masukkan Nama Startup yang ingin Dihapus (Masukkan - untuk cancel)")
		fmt.Scan(&nama)
		if nama == "-" {
			return
		}
		nama = strings.ToLower(nama)
		lokasi = sequentialSearchStartup(*tabel, nama)
		if lokasi == -1 {
			fmt.Println("Data tidak ditemukan")
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(2)
		} else {
			hapuStartup(tabel, lokasi)
			hapusAnggota(tabAnggota, lokasi)
			fmt.Println("Data Dihapus")
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(2)
		}
	case 4:
		sortIndexStartup(tabel)
		printData(*tabel)
		fmt.Println("Tekan 0 Untuk kembali")
		fmt.Scan(&pilih)
	case 5:
		return
	default:
		fmt.Println(">> Pilihan tidak valid")
		fmt.Println("Akan Kembali ke Menu Utama")
		countdownDetik(2)
	}
}

func menuKelolaTimStartup(tabAnggota *tabelAnggota, tabStartup *tabelStartup) {
	clear()
	var pilih int
	var nama string
	var lokasi int
	var banyak int
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║      🚀 KELOLA TIM STARTUP 🚀        ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. ➕ Tambah Anggota Tim              ║")
	fmt.Println("║2. 🗑️ Hapus Anggota Tim               ║")
	fmt.Println("║3. 📋 Lihat Tim per Startup           ║")
	fmt.Println("║4. 🔙 Kembali ke Menu Utama           ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		sortIndexStartup(tabStartup)
		printData(*tabStartup)
		fmt.Println("Nama Startup yang ingin ditambahkan anggotanya ? (Masukkan - untuk cancel)")
		fmt.Scan(&nama)
		if nama == "-" {
			return
		}
		nama = strings.ToLower(nama)
		lokasi = sequentialSearchStartup(*tabStartup, nama)
		if lokasi == -1 {
			fmt.Println("Data tidak ditemukan")
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(2)
		} else {
			fmt.Println("Jumlah Anggota yang Ingin Ditambah?")
			fmt.Scan(&banyak)
			inputAnggota(tabAnggota, lokasi, banyak)
			fmt.Println("Berhasil ditambahkan !")
			sortAnggota(tabAnggota)
			countdownDetik(1)
		}
		//Sandres - 103012400100
	case 2:
		sortIndexStartup(tabStartup)
		printData(*tabStartup)
		fmt.Println("Masukkan Nama Startup yang ingin Dihapus Anggotanya (Masukkan - untuk cancel)")
		fmt.Scan(&nama)
		if nama == "-" {
			return
		}
		nama = strings.ToLower(nama)
		lokasi = sequentialSearchStartup(*tabStartup, nama)
		hapusAnggota(tabAnggota, lokasi)
		fmt.Println("Data Dihapus")
		fmt.Println("Akan Kembali ke Menu Utama")
		countdownDetik(2)
	case 3:
		printData(*tabStartup)
		fmt.Println("Nama Startup yang ingin dilihat anggotanya ?")
		fmt.Scan(&nama)
		nama = strings.ToLower(nama)
		lokasi = sequentialSearchStartup(*tabStartup, nama)
		if lokasi == -1 {
			fmt.Println("Data tidak ditemukan")
			fmt.Println("Akan Kembali ke Menu Utama")
			countdownDetik(4)
		} else {
			printAnggota(*tabAnggota, lokasi, *tabStartup)
			fmt.Println("Tekan 0 Untuk kembali")
			fmt.Scan(&pilih)
			sortAnggota(tabAnggota)
		}
	case 4:
		return
	default:
		fmt.Println(">> Pilihan tidak valid")
		countdownDetik(1)
	}
}

func menuCariStartup(tabStartup *tabelStartup) {
	clear()
	var pilih int
	var dicari string
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║      🚀 PENCARIAN STARTUP 🚀         ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. 🔎 Cari Nama Startup               ║")
	fmt.Println("║2. 🔎 Cari Bidang Startup             ║")
	fmt.Println("║3. 🔙 Kembali ke Menu Utama           ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		printData(*tabStartup)
		fmt.Println("Masukkan Nama startup yang ingin di cari")
		fmt.Scan(&dicari)
		dicari = strings.ToLower(dicari)
		printKategori(*tabStartup, dicari, "nama")
		fmt.Println("Tekan 0 Untuk kembali")
		fmt.Scan(&pilih)

	case 2:
		printData(*tabStartup)
		fmt.Println("Masukkan Bidang startup yang ingin di cari")
		fmt.Scan(&dicari)
		dicari = strings.ToLower(dicari)
		printKategori(*tabStartup, dicari, "bidang")
		fmt.Println("Tekan 0 Untuk kembali")
		fmt.Scan(&pilih)

	case 3:
		return
	default:
		fmt.Println(">> Pilihan tidak valid")
		countdownDetik(1)
	}
}

func menuUrutkanStartup(tabel *tabelStartup) {
	clear()
	var pilih int
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║      🚀 PENGURUTAN STARTUP 🚀        ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║1. ⏳ Urut Tahun Berdiri              ║")
	fmt.Println("║2. 💰 Urut Pendanaan                  ║")
	fmt.Println("║3. 🔙 Kembali ke Menu Utama           ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		fmt.Println("Diurutkan Secara")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			sortTahunAscending(tabel)
			printData(*tabel)
			fmt.Println("Tekan 0 Untuk kembali")
			fmt.Scan(&pilih)
		case 2:
			sortTahunDescending(tabel)
			printData(*tabel)
			fmt.Println("Tekan 0 Untuk kembali")
			fmt.Scan(&pilih)
		}
	case 2:
		fmt.Println("Diurutkan Secara")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			sortPendanaanAscending(tabel)
			printData(*tabel)
			fmt.Println("Tekan 0 Untuk kembali")
			fmt.Scan(&pilih)
		case 2:
			sortPendanaanDescending(tabel)
			printData(*tabel)
			fmt.Println("Tekan 0 Untuk kembali")
			fmt.Scan(&pilih)
		}
	case 3:
		return
	default:
		fmt.Println(">> Pilihan tidak valid")
		countdownDetik(2)
	}
}

func menuLaporanStartup(tabel tabelStartup) {
	clear()
	var bidang [NMAX]string
	var jumlah [NMAX]int
	var totalPendanaan [NMAX]int
	var totalTahun [NMAX]int
	var n, i, j int
	jum := hitungStartUp(tabel)

	// Proses data startup
	for i = 0; i < jum; i++ {
		ditemukan := false
		j = 0
		for j < n && !ditemukan {
			if strings.ToLower(tabel[i].bidang) == strings.ToLower(bidang[j]) {
				jumlah[j]++
				totalPendanaan[j] += tabel[i].pendanaan
				totalTahun[j] += tabel[i].tahun
				ditemukan = true
			}
			j++
		}
		if !ditemukan {
			bidang[n] = tabel[i].bidang
			jumlah[n] = 1
			totalPendanaan[n] = tabel[i].pendanaan
			totalTahun[n] = tabel[i].tahun
			n++
		}
	}
	//Sandres - 103012400100
	fmt.Println("\n📌 Jumlah Startup per Bidang Usaha:")
	for i = 0; i < n; i++ {
		fmt.Printf("- %-15s : %d startup\n", bidang[i], jumlah[i])
	}

	fmt.Println("\n💰 Total Pendanaan per Bidang Usaha:")
	for i = 0; i < n; i++ {
		fmt.Printf("- %-15s : %d juta\n", bidang[i], totalPendanaan[i])
	}

	fmt.Println("\n📈 Rata-Rata Pendanaan per Bidang Usaha:")
	for i = 0; i < n; i++ {
		rata := totalPendanaan[i] / jumlah[i]
		fmt.Printf("- %-15s : %d juta\n", bidang[i], rata)
	}

	fmt.Println("\n📅 Rata-Rata Tahun Berdiri per Bidang Usaha:")
	for i = 0; i < n; i++ {
		rataTahun := totalTahun[i] / jumlah[i]
		fmt.Printf("- %-15s : %d\n", bidang[i], rataTahun)
	}

	maxJumlah := jumlah[0]
	favorit := bidang[0]
	for i = 1; i < n; i++ {
		if jumlah[i] > maxJumlah {
			maxJumlah = jumlah[i]
			favorit = bidang[i]
		}
	}
	fmt.Printf("\n🏆 Bidang usaha terfavorit: %s (%d startup)\n", favorit, maxJumlah)
}

func inputData(tabel *tabelStartup) {
	var i int
	var lanjut bool = true
	var valid bool
	var pilihan string
	for i = 0; i < NMAX && lanjut == true; i++ {
		tabel[i].idStartup = i
		fmt.Println("Masukkan Data Start Up")
		fmt.Print("Nama Start Up : ")
		fmt.Scan(&tabel[i].nama)
		fmt.Print("Bidang Start Up : ")
		fmt.Scan(&tabel[i].bidang)
		fmt.Print("Tahun Didirikan: ")
		fmt.Scan(&tabel[i].tahun)
		fmt.Print("Pendanaan (juta): ")
		fmt.Scan(&tabel[i].pendanaan)
		valid = false
		for !valid {
			fmt.Println("Tambah Data ? (Y/N)")
			fmt.Scan(&pilihan)
			pilihan = strings.ToLower(pilihan)
			if pilihan == "n" {
				lanjut = false
				valid = true
			} else if pilihan == "y" {
				lanjut = true
				valid = true
			} else {
				fmt.Println("Input tidak valid! Harap masukkan 'Y' atau 'N'.")
			}
		}

	}
}

func printData(tabel tabelStartup) { // print data startup
	clear()
	var i int
	var jum int

	jum = hitungStartUp(tabel)

	fmt.Println("╔════╦════════════════════╦════════════════════╦════════╦══════════════════╗")
	fmt.Printf("║ %-2s ║ %-18s ║ %-18s ║ %-6s ║ %-12s ║\n", "No", "Nama Startup", "Bidang Startup", "Tahun", "Pendanaan (Juta)")
	fmt.Println("╠════╬════════════════════╬════════════════════╬════════╬══════════════════╣")

	for i = 0; i < jum; i++ {
		fmt.Printf("║ %-2d ║ %-18s ║ %-18s ║ %-6d ║ %-16d ║\n",
			i+1,
			tabel[i].nama,
			tabel[i].bidang,
			tabel[i].tahun,
			tabel[i].pendanaan)
	}
	//Sandres - 103012400100
	fmt.Println("╚════╩════════════════════╩════════════════════╩════════╩══════════════════╝")
}

func inputAnggota(tabel *tabelAnggota, idStartup, jumlah int) { //input anggota berdasarkan id startup
	var i int
	var jum int = hitungAnggota(*tabel)

	for i = jum; i < jum+jumlah; i++ {
		fmt.Println("Nama Anggota, Peran Anggota")
		tabel[i].idStartupAnggota = idStartup
		fmt.Scan(&tabel[i].namaAnggota, &tabel[i].peranAnggota)
	}
}

func printAnggota(tabel tabelAnggota, idStartup int, tabStartup tabelStartup) {
	var i, no int
	var ada bool
	fmt.Println("\n=========================================")
	fmt.Printf("📋 Daftar Anggota Tim Startup #%s\n", tabStartup[idStartup].nama)
	fmt.Println("=========================================")

	fmt.Printf("%-5s %-20s %-20s\n", "No", "Nama Anggota", "Peran")
	fmt.Println("-------------------------------------------")
	//Sandres - 103012400100
	ada = false
	no = 1
	for i = 0; i < NMAX; i++ {
		if tabel[i].idStartupAnggota == idStartup && tabel[i].namaAnggota != "" {
			fmt.Printf("%-5d %-20s %-20s\n", no, tabel[i].namaAnggota, tabel[i].peranAnggota)
			no++
			ada = true
		}
	}

	if !ada {
		fmt.Println("❌ Tidak ada anggota tim yang terdaftar.")
	}
	fmt.Println()
}

func hitungAnggota(tabel tabelAnggota) int {
	var i int
	var jum int
	for i = 0; i < NMAX; i++ {
		if len(tabel[i].namaAnggota) != 0 {
			jum++
		}
	}
	return jum
}

func hitungStartUp(tabel tabelStartup) int {
	var i int
	var jum int
	for i = 0; i < NMAX; i++ {
		if len(tabel[i].nama) != 0 {
			jum++
		}
	}
	return jum
}

func sortTahunAscending(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int

	jum = hitungStartUp(*tabel)

	pass = 1

	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].tahun > tabel[i].tahun {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass = pass + 1
	}
	//Sandres - 103012400100
}

func sortTahunDescending(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int

	jum = hitungStartUp(*tabel)

	pass = 1

	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].tahun < tabel[i].tahun {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass = pass + 1
	}

}

func sortPendanaanAscending(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int

	jum = hitungStartUp(*tabel)

	pass = 1

	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].pendanaan > tabel[i].pendanaan {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass = pass + 1
	}
}

func sortPendanaanDescending(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int = hitungStartUp(*tabel)
	pass = 1

	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].pendanaan < tabel[i].pendanaan {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass++
	}
}

func sortIndexStartup(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int

	jum = hitungStartUp(*tabel)
	pass = 1
	//Sandres - 103012400100
	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].idStartup > tabel[i].idStartup {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass = pass + 1
	}
}

func sortAnggota(tabel *tabelAnggota) {
	var i, idx, pass int
	var temp anggota
	var jum int

	jum = hitungAnggota(*tabel)

	pass = 1

	for pass < jum {
		idx = pass - 1
		i = pass
		for i < jum {
			if tabel[idx].idAnggota > tabel[i].idAnggota {
				idx = i
			}
			i++
		}
		temp = tabel[pass-1]
		tabel[pass-1] = tabel[idx]
		tabel[idx] = temp
		pass = pass + 1
	}
}

func tambahStartup(tabel *tabelStartup) {
	var jum int = hitungStartUp(*tabel)

	tabel[jum].idStartup = jum
	fmt.Println("Masukkan Data Start Up")
	fmt.Print("Nama Start Up : ")
	fmt.Scan(&tabel[jum].nama)
	fmt.Print("Bidang Start Up : ")
	fmt.Scan(&tabel[jum].bidang)
	fmt.Print("Tahun Didirikan: ")
	fmt.Scan(&tabel[jum].tahun)
	fmt.Print("Pendanaan (juta): ")
	fmt.Scan(&tabel[jum].pendanaan)
}

func editStartup(tabel *tabelStartup, target int) { //target = id Startup
	fmt.Println("Masukkan Data Start Up")
	fmt.Print("Nama Start Up : ")
	fmt.Scan(&tabel[target].nama)
	fmt.Print("Bidang Start Up : ")
	fmt.Scan(&tabel[target].bidang)
	fmt.Print("Tahun Didirikan: ")
	fmt.Scan(&tabel[target].tahun)
	fmt.Print("Pendanaan (juta): ")
	fmt.Scan(&tabel[target].pendanaan)
	fmt.Println("Data berhasil diubah")

}

func hapusAnggota(tabel *tabelAnggota, idStartup int) {
	sortAnggota(tabel)
	var i int
	var jum int = hitungAnggota(*tabel)
	var idx int = 0

	for i = 0; i < jum; i++ {
		if tabel[i].idStartupAnggota != idStartup {
			(*tabel)[idx] = (*tabel)[i]
			idx++
		}
	}
	for i = idx; i < jum; i++ {
		(*tabel)[i] = anggota{}
	}
}

func sequentialSearchStartup(tabel tabelStartup, dicari string) int {
	var i int
	var jum int = hitungStartUp(tabel)
	for i = 0; i < jum; i++ {
		if strings.ToLower(tabel[i].nama) == dicari {
			return i
		}
	}
	return -1
}

func hapuStartup(tabel *tabelStartup, lokasi int) {
	var i int
	var jum int = hitungStartUp(*tabel)

	for i = lokasi; i < jum-1; i++ {
		tabel[i] = tabel[i+1]
	}

	tabel[jum-1] = startup{}
	for i = 0; i < jum-1; i++ {
		tabel[i].idStartup = i
	}
}

func printKategori(tabel tabelStartup, dicari string, bagian string) {
	clear()
	var i int
	var jum int
	var ada bool = false
	var nomor int

	jum = hitungStartUp(tabel)

	fmt.Println("╔════╦════════════════════╦════════════════════╦════════╦══════════════════╗")
	fmt.Printf("║ %-2s ║ %-18s ║ %-18s ║ %-6s ║ %-12s ║\n", "No", "Nama Startup", "Bidang Startup", "Tahun", "Pendanaan (Juta)")
	fmt.Println("╠════╬════════════════════╬════════════════════╬════════╬══════════════════╣")
	if bagian == "nama" {
		for i = 0; i < jum; i++ {
			if strings.ToLower(tabel[i].nama) == dicari {
				fmt.Printf("║ %-2d ║ %-18s ║ %-18s ║ %-6d ║ %-16d ║\n",
					nomor+1,
					tabel[i].nama,
					tabel[i].bidang,
					tabel[i].tahun,
					tabel[i].pendanaan)
				ada = true
				nomor++
			}
			//Sandres - 103012400100
		}
	} else if bagian == "bidang" {
		for i = 0; i < jum; i++ {
			if strings.ToLower(tabel[i].bidang) == dicari {
				fmt.Printf("║ %-2d ║ %-18s ║ %-18s ║ %-6d ║ %-16d ║\n",
					nomor+1,
					tabel[i].nama,
					tabel[i].bidang,
					tabel[i].tahun,
					tabel[i].pendanaan)
				ada = true
				nomor++
			}
		}

	}
	fmt.Println("╚════╩════════════════════╩════════════════════╩════════╩══════════════════╝")
	if !ada {
		fmt.Println("Data tidak ditemukan")
	}
}

func main() {
	var menu, pilihan int
	var tabStartup tabelStartup
	var tabAnggota tabelAnggota
	var menu1 bool = true
	var role int
	var pilih int
	var nama string
	var lokasi int
	menu = 1
	for menu != 0 && menu1 {
		clear()
		menuAwal()
		fmt.Scan(&menu)
		switch menu {
		case 1:
			inputData(&tabStartup)
			menu1 = false
		case 2:
			loadDataStartup(&tabStartup)
			fmt.Println("Data loaded...")
			countdownDetik(1)
			menu1 = false
		case 3:
			changeLog()
		case 4:
			os.Exit(0)
		default:
			clear()
			menuAwal()
			fmt.Scan(&menu)
		}

	}
	for {
		clear()
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║      ⚙️     Masuk Sebagai      ⚙️    ║")
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║1. 🔧 Admin                           ║")
		fmt.Println("║2. 👤 Pengguna                        ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Println("Masukkan peran anda")
		fmt.Scan(&role)
		for menu != 0 && role == 1 {
			tampilkanMenuUtama()
			fmt.Println("ADMIN 🔧")
			fmt.Print("Pilih menu (1-6): ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				menuKelolaDataStartup(&tabStartup, &tabAnggota)
			case 2:
				menuKelolaTimStartup(&tabAnggota, &tabStartup)
			case 3:
				menuCariStartup(&tabStartup)
			case 4:
				menuUrutkanStartup(&tabStartup)
			case 5:
				menuLaporanStartup(tabStartup)
				fmt.Scan(&pilihan)
			case 6:
				fmt.Println("\nTerima kasih telah menggunakan Aplikasi Manajemen Startup. Sampai Jumpa! 👋")
				os.Exit(0)
			case 9:
				role = 0
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
				countdownDetik(2)
			}
		}
		// Sandres - 103012400100
		for menu != 0 && role == 2 {
			menuPengguna()
			fmt.Println("PENGGUNA 👤")
			fmt.Print("Pilih menu (1-6): ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				sortIndexStartup(&tabStartup)
				printData(tabStartup)
				fmt.Println("Tekan 0 Untuk kembali")
				fmt.Scan(&pilih)
			case 2:
				sortIndexStartup(&tabStartup)
				printData(tabStartup)
				fmt.Println("Nama Startup yang ingin dilihat anggotanya ?")
				fmt.Scan(&nama)
				nama = strings.ToLower(nama)
				lokasi = sequentialSearchStartup(tabStartup, nama)
				if lokasi == -1 {
					fmt.Println("Data tidak ditemukan")
					fmt.Println("Akan Kembali ke Menu Utama")
					countdownDetik(4)
				} else {
					printAnggota(tabAnggota, lokasi, tabStartup)
					fmt.Println("Tekan 0 Untuk kembali")
					fmt.Scan(&pilih)
					sortAnggota(&tabAnggota)
				}
			case 3:
				menuCariStartup(&tabStartup)
			case 4:
				menuUrutkanStartup(&tabStartup)
			case 5:
				menuLaporanStartup(tabStartup)
				fmt.Scan(&pilihan)
			case 6:
				fmt.Println("\nTerima kasih telah menggunakan Aplikasi Manajemen Startup. Sampai Jumpa! 👋")
				os.Exit(0)
			case 9:
				role = 0
			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
				countdownDetik(2)
			}
		}
	}
}
func loadDataStartup(tabel *tabelStartup) {
	data := []startup{
		{nama: "Luminous", bidang: "Pendidikan", tahun: 2023, pendanaan: 750},
		{nama: "FlowState", bidang: "Kesehatan", tahun: 2021, pendanaan: 1250},
		{nama: "AgriGrow", bidang: "Pertanian", tahun: 2024, pendanaan: 500},
		{nama: "SwiftLogix", bidang: "Logistik", tahun: 2022, pendanaan: 900},
		{nama: "Artify", bidang: "Kreator", tahun: 2020, pendanaan: 1500},
		{nama: "MediSync", bidang: "Kesehatan", tahun: 2019, pendanaan: 2000},
		{nama: "CodeNest", bidang: "Teknologi", tahun: 2021, pendanaan: 1100},
		{nama: "EcoTrack", bidang: "Lingkungan", tahun: 2022, pendanaan: 850},
		{nama: "SmartAgro", bidang: "Pertanian", tahun: 2023, pendanaan: 950},
		{nama: "SkillForge", bidang: "Pendidikan", tahun: 2020, pendanaan: 1200},
		{nama: "HomeEase", bidang: "Properti", tahun: 2021, pendanaan: 600},
		{nama: "QuickPay", bidang: "Fintech", tahun: 2024, pendanaan: 1400},
		{nama: "TransPort", bidang: "Transportasi", tahun: 2020, pendanaan: 700},
		{nama: "FoodWise", bidang: "Kuliner", tahun: 2023, pendanaan: 950},
		{nama: "BuildNow", bidang: "Konstruksi", tahun: 2018, pendanaan: 1800},
		{nama: "InnoSpace", bidang: "Teknologi", tahun: 2021, pendanaan: 2100},
		{nama: "Recyco", bidang: "Lingkungan", tahun: 2022, pendanaan: 770},
		{nama: "WellNest", bidang: "Kesehatan", tahun: 2020, pendanaan: 1600},
		{nama: "DigiTani", bidang: "Pertanian", tahun: 2024, pendanaan: 880},
		{nama: "Craftory", bidang: "Kreator", tahun: 2023, pendanaan: 1020},
		{nama: "FundFlow", bidang: "Fintech", tahun: 2022, pendanaan: 1300},
		{nama: "SmartStay", bidang: "Properti", tahun: 2021, pendanaan: 990},
	}

	for i, d := range data {
		tabel[i] = d
		tabel[i].idStartup = i
	}
}

func changeLog() {
	clear()
	var pilih int
	fmt.Println("v0.0.0 membuat func main serta tampilan awal")
	fmt.Println("v0.1.1 menambahkan fitur sehingga pengguna dapat berinteraski dengan aplikasi")
	fmt.Println("v0.2.2 membuat fungsi bertujuan untuk memasukkan data startup dan print data startup")
	fmt.Println("v0.3.3 memperbaiki tampilan data, dan menambahkan fitur")
	fmt.Println("v0.4.4 perbaikan fitur untuk menampilkan data")
	fmt.Println("v0.5.5 menambahkan fungsi baru untuk mencari data")
	fmt.Println("v0.6.6 perbaikan opsi menu, perbaikan bug tampilan, perbaikan fitur")
	fmt.Println("v0.7.1 menambahkan fitur user dan admin")
	fmt.Println("v0.7.7 memperbaiki bug penambahan data startup")
	fmt.Println()
	fmt.Println("Tekan 0 Untuk kembali")
	fmt.Scan(&pilih)
	if pilih == 0 {
		return
	}
}
