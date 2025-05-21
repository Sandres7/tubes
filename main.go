package main

import "fmt"

const NMAX int = 1000

//Menampung macam macam data pada startup
type startup struct {
	idStartup int
	nama      string //Kelipatan 10
	pendanaan int
	tahun     int
	bidang    string
}

//Menampung macam macam data yang id miliki anggota
type anggota struct {
	idStartupAnggota int
	idAnggota        int
	namaAnggota      string
	peranAnggota     string
}

type tabelAnggota [NMAX]anggota //inisiasi arrray anggota
type tabelStartup [NMAX]startup //inisiasi array startup

func printMenu() { // opsi menu pilihan utama
	fmt.Println()
	fmt.Println("----------------------------------------------")
	fmt.Println("Selamat datang di layanan Start Up (beta)")
	fmt.Println("Silakan pilih menu :")
	fmt.Println("1. Masukkan data startup")
	fmt.Println("2. Print data startUp")
	fmt.Println("3. Masukkan data Anggota")
	fmt.Println("4. Print Data Anggota")
	fmt.Println("5. Urutkan Startup")
	fmt.Println("6. Hapus Data StartUp")
	fmt.Println("7. Tambah Startup Baru")
	fmt.Println("----------------------------------------------")

}

func menuPilihanPengurutan() {
	fmt.Println("Pilih yang ingin diurutkan")
	fmt.Println("1. Berdasarkan Tahun")
	fmt.Println("2. Berdasarkan Pendanaan")
}

func menuPengurutan() {
	fmt.Println("Diurutkan Secara")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
}

//Nama startUp Bidang Tahun Pendanaan
func inputData(tabel *tabelStartup) {
	var i int
	var lanjut bool = true
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
		fmt.Println("Tambah Data ? (Y/N)")
		fmt.Scan(&pilihan)
		if pilihan == "N" {
			lanjut = false
		}

	}
}

func printData(tabel tabelStartup) { // print data startup
	var i int
	var jum int

	jum = hitungStartUp(tabel)
	fmt.Println("Id StartUp \t Nama StartUP: \t Bidang StartUp \t Tahun Didirikan \t Pendanaan (Juta): ")
	for i = 0; i < jum; i++ {
		fmt.Printf("%d. %d %s \t %s \t\t %d \t\t\t %d \n", i+1, tabel[i].idStartup, tabel[i].nama, tabel[i].bidang, tabel[i].tahun, tabel[i].pendanaan)
	}
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

func printAnggota(tabel tabelAnggota, idStartup int) {
	var i int
	idStartup--
	for i = 0; i < NMAX; i++ {
		if len(tabel[i].namaAnggota) != 0 && tabel[i].idStartupAnggota == idStartup {
			fmt.Println(tabel[i].idStartupAnggota, tabel[i].namaAnggota, tabel[i].peranAnggota)
		}
	}
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

func sortIndexStartup(tabel *tabelStartup) {
	var i, idx, pass int
	var temp startup
	var jum int

	jum = hitungStartUp(*tabel)
	pass = 1

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
	var i, idBaru int
	var stop bool = false
	var jum int = hitungStartUp(*tabel)
	for i = 0; i < jum && stop == false; i++ {
		if tabel[i+1].idStartup-tabel[i].idStartup > 1 {
			idBaru = i + 1
			stop = true
		} else {
			idBaru = jum
		}
	}
	tabel[jum].idStartup = idBaru
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
	target--
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
	var i, tot, j int
	var jum int = hitungAnggota(*tabel)
	for i = 0; i < jum; i++ {
		if tabel[i].idStartupAnggota == idStartup {
			tot++
		}
	}

	for i = 0; i < jum; i++ {
		if tabel[i].idStartupAnggota == idStartup {
			for j = i; j < jum; j++ {
				tabel[j] = tabel[j+1]
			}
			jum--
		}

	}
	for i = 0; i < tot; i++ {
		tabel[jum+i] = anggota{}
	}
}

func binSearch(tabel tabelStartup, dicari int) int {
	var left, right, mid int
	var idx int
	var jum int
	sortIndexStartup(&tabel)
	jum = hitungStartUp(tabel)

	left = 0
	right = jum - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if dicari < tabel[mid].idStartup {
			right = mid - 1
		} else if dicari > tabel[mid].idStartup {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func hapuStartup(tabel *tabelStartup, lokasi int) {
	var i int
	var jum int = hitungStartUp(*tabel)
	sortIndexStartup(tabel)
	for i = lokasi; i < jum; i++ {
		tabel[i] = tabel[i+1]
	}
	tabel[jum] = startup{} //opsional karena array panjang
}

func main() {
	var tabStartup tabelStartup
	var tabAnggota tabelAnggota
	var menu int
	var pilihanData int
	var jumlahAnggota int
	printMenu()
	fmt.Scan(&menu)
	for menu != 0 {
		switch menu {
		case 1:
			inputData(&tabStartup)
		case 2:
			printData(tabStartup)
		case 3:
			fmt.Println("Anggota Kelompok yang ingin di ubah")
			printData(tabStartup)
			fmt.Scan(&pilihanData)
			fmt.Println("Jumlah Anggota")
			fmt.Scan(&jumlahAnggota)
			inputAnggota(&tabAnggota, pilihanData, jumlahAnggota)
		case 4:
			fmt.Println("Anggota Kelompok yang ingin di Tampilkan")
			printData(tabStartup)
			fmt.Scan(&pilihanData)
			printAnggota(tabAnggota, pilihanData)
		case 5:
			menuPilihanPengurutan()
			fmt.Scan(&pilihanData)
			switch pilihanData {
			case 1:
				menuPengurutan()
				fmt.Scan(&pilihanData)
				switch pilihanData {
				case 1:
					sortTahunAscending(&tabStartup)
					fmt.Println("Ascending Tahun")
					printData(tabStartup)
				case 2:
					sortTahunDescending(&tabStartup)
					fmt.Println("Descending Tahun")
					printData(tabStartup)
				}
			case 2:
				menuPengurutan()
				fmt.Scan(&pilihanData)
				switch pilihanData {
				case 1:
					sortPendanaanAscending(&tabStartup)
					fmt.Println("Ascending Pendanaan")
					printData(tabStartup)
				case 2:
					sortPendanaanDescending(&tabStartup)
					fmt.Println("Descending Pendanaan")
					printData(tabStartup)
				}
			}
		case 6:
			fmt.Println("Pilih ID startup yang ingin di Hapus")
			fmt.Scan(&pilihanData)
			pilihanData = binSearch(tabStartup, pilihanData)
			if pilihanData == -1 {
				fmt.Println("Startup tidak di temukan.")
			} else {
				hapuStartup(&tabStartup, pilihanData)
				printData(tabStartup)
			}
			hapusAnggota(&tabAnggota, pilihanData)
		case 7:
			fmt.Println("Startup baru yang ingin di tambahkan")
			sortIndexStartup(&tabStartup)
			tambahStartup(&tabStartup)
			sortIndexStartup(&tabStartup)
			fmt.Println("Data Berhasil Di tambahkan")
		}
		printMenu()
		fmt.Scan(&menu)
	}
}
