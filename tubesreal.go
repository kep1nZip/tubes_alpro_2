package main

import "fmt"

const TMAX int = 5
const PMAX int = 5

type team struct {
	nama     string
	pemain   [PMAX]player
	menang   int
	selisih  int
	totalACS int
}

type player struct {
	nama string
	ACS  int
}

var tournaments [TMAX]team
var totalTim int
var t team

func main() {
	var n int

	for {
		fmt.Println("\n|====================================================|")
		fmt.Println("|    APLIKASI PENGELOLAAN DATA E-Sport TOURNAMENT    |")
		fmt.Println("|====================================================|")
		fmt.Println("|                  >{===MENU===}<                    |")
		fmt.Println("|----------------------------------------------------|")
		fmt.Println("| 1. Tambah Tim + player                             |")
		fmt.Println("| 2. Edit Tim                                        |")
		fmt.Println("| 3. Hapus Tim                                       |")
		fmt.Println("| 4. Urutkan berdasarkan Menang (Selection Sort)     |")
		fmt.Println("| 5. Urutkan berdasarkan Menang (Insertion Sort)     |")
		fmt.Println("| 6. Cari Tim (Sequential Search)                    |")
		fmt.Println("| 7. Cari Tim (Binary Search)                        |")
		fmt.Println("| 8. Tampilkan Statistik                             |")
		fmt.Println("| 9. Update Klasemen                                 |")
		fmt.Println("| 0. Keluar                                          |")
		fmt.Println("|====================================================|")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			addTeam(&tournaments, &totalTim)
		case 2:
			editTeam(&tournaments, totalTim)
		case 3:
			deleteTeam(&tournaments, &totalTim)
		case 4:
			winSelection(&tournaments, totalTim)
			print(tournaments, totalTim)
		case 5:
			winInsertion(&tournaments, totalTim)
			print(tournaments, totalTim)
		case 6:
			sequentialSearch(tournaments, &totalTim)
		case 7:
			binarySearch(tournaments, &totalTim)
		case 8:
			showStats(tournaments, totalTim)
		case 9:
			updateKlasemen(&tournaments, totalTim)
		case 0:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Masukan Kembali Pilihan Yang Tertera")
		}
	}
}

// TAMBAH TEAM
func addTeam(tournaments *[TMAX]team, totalTim *int) {
	var t team
	var i int
	var confirm string
	var valid bool

	if *totalTim >= TMAX {
		fmt.Println("Kapasitas tim penuh!")
		return
	}
	t.totalACS = 0

	fmt.Println("|===========TAMBAH TIM============|")

	fmt.Println()
	fmt.Print("Masukkan Nama Tim (Pakai '_' untuk spasi, Contoh: Apple_Team): ")
	fmt.Scan(&t.nama)

	fmt.Println()
	fmt.Println("|==========TAMBAH PEMAIN==========|")
	fmt.Println()

	fmt.Println("Masukkan data 5 player:")
	for i = 0; i < 5; i++ {
		fmt.Printf("Player %d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&t.pemain[i].nama)

		valid = false
		for !valid {
			fmt.Print("Total ACS (0 - 400): ")
			fmt.Scan(&t.pemain[i].ACS)

			if t.pemain[i].ACS < 0 || t.pemain[i].ACS > 400 {
				fmt.Println("Nilai ACS harus di antara 0 hingga 400, Tolong Input Kembali.")
			} else {
				valid = true
			}
		}
		t.totalACS += t.pemain[i].ACS

	}
	valid = false
	for !valid {
		fmt.Print("Masukkan jumlah kemenangan tim: ")
		fmt.Scan(&t.menang)
		if t.menang < 0 {
			fmt.Println("Input Tidak Boleh Kurang Dari 0, Tolong Masukan Input Kembali")
		} else {
			valid = true
		}
	}

	fmt.Print("Masukkan selisih skor tim: ")
	fmt.Scan(&t.selisih)

	valid = false
	for !valid {
		fmt.Printf("Konfirmasi Untuk Menambahkan Tim '%s'? (y/n): ", t.nama)
		fmt.Scan(&confirm)
		if confirm == "n" || confirm == "N" {
			fmt.Println("Data Tidak Disimpan")
			valid = true
		} else if confirm == "y" || confirm == "Y" {
			valid = true
			tournaments[*totalTim] = t
			*totalTim++
			fmt.Println("Tim beserta 5 pemain berhasil ditambahkan!")
		} else {
			fmt.Println("Input Tidak Valid, Tolong Masukan Input Yang Diminta")
		}
	}

}

//cetak
func print(tournaments [TMAX]team, totalteam int) {
	var i int
	if totalTim == 0 {
		return
	}
	fmt.Print("==========================================================================\n")
	fmt.Printf("| %-15s | %-15s | %-15s | %-15s | \n", "Nama Tim", "Menang", "Selisih", "Total ACS")
	fmt.Print("--------------------------------------------------------------------------\n")
	for i = 0; i < totalteam; i++ {
		fmt.Printf("|%d. %-15s | %-15d | %-15d | %-15d |\n", i+1, tournaments[i].nama, tournaments[i].menang, tournaments[i].selisih, tournaments[i].totalACS)
	}
	fmt.Print("==========================================================================")

}

// EDIT TEAM
func editTeam(tournaments *[TMAX]team, totalTim int) {
	var ketemu bool
	var i, j int
	var playerBaru, timBaru, namaTeam string
	var ACSUpdate, menangBaru, selisihBaru int

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	fmt.Print("===Nama Tim Yang Terdaftar===\n")
	for i = 0; i < totalTim; i++ {
		fmt.Printf("%d. %s\n", i+1, tournaments[i].nama)
	}
	fmt.Print("=============================")
	fmt.Println()

	fmt.Print("Masukkan nama tim yang akan diedit: ")
	fmt.Scan(&namaTeam)

	ketemu = false
	i = 0
	for i < totalTim && !ketemu {
		if tournaments[i].nama == namaTeam {
			fmt.Println("\nEdit data tim:")
			fmt.Print("Nama tim baru (isi '-' jika tidak ingin mengganti): ")
			fmt.Scan(&timBaru)

			if timBaru != "-" {
				tournaments[i].nama = timBaru
			}

			fmt.Println("Edit data player: ")
			for j = 0; j < 5; j++ {
				fmt.Printf("Player %d (%s): \n", j+1, tournaments[i].pemain[j].nama)
				fmt.Print("  Nama baru (isi '-' jika tidak ingin mengubah): ")
				fmt.Scan(&playerBaru)
				if playerBaru != "-" {
					tournaments[i].pemain[j].nama = playerBaru
				}

				fmt.Print("Jumlah ACS baru (0 jika tidak ingin mengubah): ")
				fmt.Scan(&ACSUpdate)
				if ACSUpdate != 0 {
					tournaments[i].pemain[j].ACS = ACSUpdate
				}
			}

			fmt.Print("Jumlah kemenangan baru (0 jika tidak ingin mengubah): ")
			fmt.Scan(&menangBaru)
			if menangBaru != 0 {
				tournaments[i].menang = menangBaru
			}
			fmt.Print("Selisih skor baru (0 jika tidak ingin mengubah): ")
			fmt.Scan(&selisihBaru)
			if selisihBaru != 0 {
				tournaments[i].selisih = selisihBaru
			}
			fmt.Println("Tim berhasil diperbarui!")
			ketemu = true
		}
		i++
	}

	if !ketemu {
		fmt.Println("Tim tidak ditemukan!")
	}
}

// DELETE TEAM
func deleteTeam(tournaments *[TMAX]team, totalTim *int) {
	var found bool
	var i, j int
	var namaTeam string

	if *totalTim == 0 {
		fmt.Println("Belum ada tim!")
		return
	}

	fmt.Print("===Nama Tim Yang Terdaftar===\n")
	for i = 0; i < *totalTim; i++ {
		fmt.Printf("%d. %s\n", i+1, tournaments[i].nama)
	}
	fmt.Print("=============================")
	fmt.Println()
	fmt.Print("Masukkan nama tim yang akan dihapus: ")
	fmt.Scan(&namaTeam)

	found = false
	for i = 0; i < *totalTim; i++ {
		if tournaments[i].nama == namaTeam {
			for j = i; j < *totalTim-1; j++ {
				tournaments[j] = tournaments[j+1]
			}
			*totalTim = *totalTim - 1
			fmt.Println("Tim berhasil dihapus!")
			found = true
			i = *totalTim
		}
	}

	if !found {
		fmt.Println("Tim tidak ditemukan!")
	}
}

// SEQUENTIAL SEARCH CARI NAMA TIM
func sequentialSearch(tournaments [TMAX]team, totalTim *int) {
	var ketemu int
	var k, i int
	var namaTeam string

	if *totalTim == 0 {
		fmt.Println("Belum ada tim!")
		return
	}

	fmt.Print("===Nama Tim Yang Terdaftar===\n")
	for i = 0; i < *totalTim; i++ {
		fmt.Printf("%d. %s\n", i+1, tournaments[i].nama)
	}
	fmt.Print("=============================")
	fmt.Println()

	fmt.Print("Masukkan nama tim yang dicari: ")
	fmt.Scan(&namaTeam)

	ketemu = -1
	k = 0
	for ketemu == -1 && k < *totalTim {
		if tournaments[k].nama == namaTeam {
			ketemu = k
		}
		k++
	}

	if ketemu != -1 {
		fmt.Println("Tim yang dicari KETEMU! :")
		fmt.Println("Nama Tim:", tournaments[ketemu].nama)
		fmt.Println("Jumlah Menang:", tournaments[ketemu].menang)
		fmt.Println("Jumlah Selisih Skor:", tournaments[ketemu].selisih)
		fmt.Println("Daftar Player:")
		for i = 0; i < PMAX; i++ {
			fmt.Printf("- %s (ACS: %d)\n", tournaments[ketemu].pemain[i].nama, tournaments[ketemu].pemain[i].ACS)
		}
	} else {
		fmt.Println("\nTim dengan nama", namaTeam, "tidak terdaftar")
	}
}

// BINARY SEARCH CARI NAMA TIM
func binarySearch(tournaments [TMAX]team, totalTim *int) {
	var i, j int
	var left, right, mid, idx int
	var namaTeam string
	var temp team
	var sort [TMAX]team
	var ketemu bool

	if *totalTim == 0 {
		fmt.Println("Belum ada tim!")
		return
	}

	for i = 0; i < *totalTim; i++ {
		sort[i] = tournaments[i]
	}

	for i = 1; i < *totalTim; i++ {
		temp = sort[i]
		j = i - 1
		for j >= 0 && sort[j].nama > temp.nama {
			sort[j+1] = sort[j]
			j--
		}
		sort[j+1] = temp
	}

	fmt.Print("===Nama Tim Yang Terdaftar===\n")
	for i = 0; i < *totalTim; i++ {
		fmt.Printf("%d. %s\n", i+1, tournaments[i].nama)
	}
	fmt.Print("=============================")
	fmt.Println()

	fmt.Print("Masukkan nama tim yang dicari: ")
	fmt.Scan(&namaTeam)

	left = 0
	right = *totalTim - 1
	ketemu = false
	idx = -1

	for left <= right && !ketemu {
		mid = (left + right) / 2

		if sort[mid].nama == namaTeam {
			ketemu = true
			i = 0
			for i < *totalTim && idx == -1 {
				if tournaments[i].nama == namaTeam {
					idx = i
				}
				i++
			}
		} else if sort[mid].nama < namaTeam {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if ketemu {
		fmt.Println("Tim yang dicari KETEMU! :")
		fmt.Println("Nama Tim:", tournaments[idx].nama)
		fmt.Println("Jumlah Menang:", tournaments[idx].menang)
		fmt.Println("Jumlah Selisih Skor:", tournaments[idx].selisih)
		fmt.Println("Daftar Player:")
		for j = 0; j < PMAX; j++ {
			fmt.Printf("- %s (ACS: %d)\n", tournaments[idx].pemain[j].nama, tournaments[idx].pemain[j].ACS)
		}
		return
	} else {
		fmt.Printf("\nTim dengan nama %s tidak terdaftar\n", namaTeam)
	}
}

// KLASEMEN
func updateKlasemen(tournaments *[TMAX]team, totalteam int) {
	var namaTeam string
	var menangBaru, selisihBaru int
	var found bool
	var i int

	if totalTim == 0 {
		fmt.Println("Belum ada data!")
		return
	}

	fmt.Print("===Nama Tim Yang Terdaftar===\n")
	for i = 0; i < totalTim; i++ {
		fmt.Printf("%d. %s\n", i+1, tournaments[i].nama)
	}
	fmt.Print("=============================")
	fmt.Println()

	fmt.Print("Masukkan nama tim yang dicari: ")
	fmt.Scan(&namaTeam)

	found = false
	i = 0
	for i < totalTim && !found {
		if tournaments[i].nama == namaTeam {
			fmt.Printf("Data saat ini:\nMenang: %d\nSelisih: %d\n", tournaments[i].menang, tournaments[i].selisih)
			fmt.Print("Tambah Jumlah kemenangan: ")
			fmt.Scan(&menangBaru)
			fmt.Print("Tambah selisih skor: ")
			fmt.Scan(&selisihBaru)

			tournaments[i].menang = tournaments[i].menang + menangBaru
			tournaments[i].selisih = tournaments[i].selisih + selisihBaru

			fmt.Println("Klasemen diperbarui")
			fmt.Printf("Data baru:\nMenang: %d\nSelisih: %d\n", tournaments[i].menang, tournaments[i].selisih)
			found = true
		}
		i++
	}
	if !found {
		fmt.Println("Tim tidak ada")
	}
}

// menang disorting menggunakan Selection sort
func winSelection(tournaments *[TMAX]team, totalteam int) {
	var i, idx, pass int
	var temp team
	var req1, req2, req3 bool

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	pass = 1
	for pass < totalteam {
		idx = pass - 1
		i = pass

		for i < totalteam {
			req1 = tournaments[i].menang > tournaments[idx].menang
			req2 = tournaments[i].menang == tournaments[idx].menang && tournaments[i].selisih > tournaments[idx].selisih
			req3 = (tournaments[i].menang == tournaments[idx].menang && tournaments[i].selisih == tournaments[idx].selisih) && tournaments[i].totalACS > tournaments[idx].totalACS
			if req1 || req2 || req3 {
				idx = i
			}
			i = i + 1
		}
		temp = tournaments[pass-1]
		tournaments[pass-1] = tournaments[idx]
		tournaments[idx] = temp
		pass = pass + 1
	}
}

// menang disorting menggunakan Insertion sort
func winInsertion(tournaments *[TMAX]team, totalTim int) {
	var i, pass int
	var temp team
	var req1, req2, req3 bool

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	pass = 1

	for pass < totalTim {
		i = pass
		temp = tournaments[pass]
		req1 = tournaments[i-1].menang < temp.menang
		req2 = tournaments[i-1].menang == temp.menang && tournaments[i-1].selisih < temp.selisih
		req3 = (tournaments[i-1].menang == temp.menang && tournaments[i-1].selisih == temp.selisih) && tournaments[i-1].totalACS < temp.totalACS

		for i > 0 && req1 || req2 || req3 {
			tournaments[i] = tournaments[i-1]
			i = i - 1
		}

		tournaments[i] = temp
		pass = pass + 1
	}
}

//menampilkan statistik
func showStats(tournaments [TMAX]team, totalTim int) {
	var MVP, p player
	var i, j int
	var MVPteam team

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	MVP.ACS = -1
	for i = 0; i < totalTim; i++ {
		fmt.Printf("Nama Tim: %s\n", tournaments[i].nama)
		fmt.Printf("Total Menang: %d\n", tournaments[i].menang)
		fmt.Printf("Total Selisih: %d\n", tournaments[i].selisih)
		fmt.Printf("Total ACS Tim: %d\n", tournaments[i].totalACS)

		fmt.Println()

		for j = 0; j < PMAX; j++ {
			fmt.Printf("Nama Pemain: %s\n", tournaments[i].pemain[j].nama)
			fmt.Printf("Total ACS (Average Combat Score): %d\n", tournaments[i].pemain[j].ACS)
			fmt.Println()
			p = tournaments[i].pemain[j]
			if p.ACS > MVP.ACS {
				MVP = p
				MVPteam = tournaments[i]
			} else if p.ACS == MVP.ACS {
				if tournaments[i].menang > MVPteam.menang {
					MVP = p
					MVPteam = tournaments[i]
				} else if tournaments[i].menang == MVPteam.menang {
					if tournaments[i].selisih > MVPteam.selisih {
						MVP = p
						MVPteam = tournaments[i]
					}
				} else if (tournaments[i].menang == MVPteam.menang) && (tournaments[i].selisih == MVPteam.selisih) {
					if tournaments[i].totalACS > MVPteam.totalACS {
						MVP = p
						MVPteam = tournaments[i]
					}
				}
			}
		}

	}

	fmt.Print("===Pemain Terbaik===\n")
	fmt.Printf("Pemain : %s\n", MVP.nama)
	fmt.Printf("ACS    : %d\n", MVP.ACS)

}
