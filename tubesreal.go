package main

import "fmt"

const TMAX int = 10
const PMAX int = 5

type team struct {
	nama    string
	pemain  [PMAX]player
	menang  int
	selisih int
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
	var namaTeam string

	for {
		fmt.Println("\n== Menu ==")
		fmt.Println("1. Tambah Tim + player")
		fmt.Println("2. Edit Tim")
		fmt.Println("3. Hapus Tim")
		fmt.Println("4. Urutkan berdasarkan Menang (Selection Sort)")
		fmt.Println("5. Urutkan berdasarkan Menang (Insertion Sort)")
		fmt.Println("6. Cari Tim (Sequential Search)")
		fmt.Println("7. Cari Tim (Binary Search)")
		fmt.Println("8. Tampilkan Statistik")
		fmt.Println("9. Update Klasemen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&n)

		switch n {
		case 1:
			addTeam(&tournaments, &totalTim)
		case 2:
			fmt.Print("Masukkan nama tim yang akan diedit: ")
			fmt.Scan(&namaTeam)
			editTeam(&tournaments, totalTim, namaTeam)
		case 3:
			fmt.Print("Masukkan nama tim yang akan dihapus: ")
			fmt.Scan(&namaTeam)
			deleteTeam(&tournaments, &totalTim, namaTeam)
		case 4:
			winSelection(&tournaments, totalTim)
			print(tournaments, totalTim)
		case 5:
			winInsertion(&tournaments, totalTim)
			print(tournaments, totalTim)
		case 6:
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&namaTeam)
			sequentialSearch(tournaments, totalTim, namaTeam)
		case 7:
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&namaTeam)
			binarySearch(tournaments, totalTim, namaTeam)
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

	if *totalTim >= TMAX {
		fmt.Println("Kapasitas tim penuh!")
		return
	}

	fmt.Print("Masukkan nama tim: ")
	fmt.Scan(&t.nama)

	fmt.Println("Masukkan data 5 player:")
	for i = 0; i < 5; i++ {
		fmt.Printf("Player %d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&t.pemain[i].nama)
		fmt.Print("Total ACS: ")
		fmt.Println("0 - 400")
		fmt.Scan(&t.pemain[i].ACS)
	}

	fmt.Print("Masukkan jumlah kemenangan tim: ")
	fmt.Scan(&t.menang)
	fmt.Print("Masukkan selisih skor tim: ")
	fmt.Scan(&t.selisih)

	tournaments[*totalTim] = t
	*totalTim++
	fmt.Println("Tim beserta 5 pemain berhasil ditambahkan!")
}

//cetak
func print(tournaments [TMAX]team, totalteam int) {
	var i int
	if totalTim == 0 {
		return
	}

	fmt.Printf("| %-15s | %-15s | %-15s |\n", "Nama Tim", "Menang", "Selisih")
	fmt.Print("-------------------------------------------------------\n")
	for i = 0; i < totalteam; i++ {
		fmt.Printf("| %-15s | %-15d | %-15d |\n", tournaments[i].nama, tournaments[i].menang, tournaments[i].selisih)
	}

}

// EDIT TEAM
func editTeam(tournaments *[TMAX]team, totalTim int, namaTeam string) {
	var ketemu bool
	var i, j int
	var playerBaru, timBaru string
	var ACSUpdate, menangBaru, selisihBaru int

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

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
func deleteTeam(tournaments *[TMAX]team, totalTim *int, teamName string) {
	var found bool
	var i, j int

	if *totalTim == 0 {
		fmt.Println("Belum ada tim!")
		return
	}

	found = false
	for i = 0; i < *totalTim; i++ {
		if tournaments[i].nama == teamName {
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
func sequentialSearch(tournaments [TMAX]team, totalTim int, namaTeam string) {
	var ketemu int
	var k, i int

	ketemu = -1
	k = 0
	for ketemu == -1 && k < totalTim {
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
func binarySearch(tournaments [TMAX]team, totalTim int, namaTeam string) {
	var i int
	var left, right, mid, hasil int
	var ketemu bool

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	ketemu = false
	left = 0
	right = totalTim - 1
	mid = (left + right) / 2

	for left <= right && !ketemu {
		if tournaments[mid].nama == namaTeam {
			ketemu = true
			hasil = mid
		} else if tournaments[mid].nama < namaTeam {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if ketemu {
		fmt.Println("Tim yang dicari KETEMU! :")
		fmt.Println("Nama Tim:", tournaments[hasil].nama)
		fmt.Println("Jumlah Menang:", tournaments[hasil].menang)
		fmt.Println("Jumlah Selisih Skor:", tournaments[hasil].selisih)
		fmt.Println("Daftar Player:")
		for i = 0; i < PMAX; i++ {
			fmt.Printf("- %s (ACS: %d)\n", tournaments[hasil].pemain[i].nama, tournaments[hasil].pemain[i].ACS)
		}
	} else {
		fmt.Println("\nTim dengan nama", namaTeam, "tidak terdaftar")
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

	fmt.Print("Nama tim: ")
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

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	pass = 1
	for pass < totalteam {
		idx = pass - 1
		i = pass
		for i < totalteam {
			if tournaments[i].menang > tournaments[idx].menang {
				idx = i
			} else if tournaments[i].menang == tournaments[idx].menang {
				if tournaments[i].selisih > tournaments[idx].selisih {
					idx = i
				}
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
	var req1, req2 bool

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

		for i > 0 && req1 || req2 {
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

	if totalTim == 0 {
		fmt.Println("Belum ada data tim!")
		return
	}

	MVP.ACS = -1
	for i = 0; i < totalTim; i++ {
		fmt.Printf("Nama Tim: %s\n", tournaments[i].nama)
		fmt.Printf("Total Menang: %d\n", tournaments[i].menang)
		fmt.Printf("Total Selisih: %d\n", tournaments[i].selisih)
		fmt.Println()

		for j = 0; j < PMAX; j++ {
			fmt.Printf("Nama Pemain: %s\n", tournaments[i].pemain[j].nama)
			fmt.Printf("Total ACS (Avarage Combat Score): %d\n", tournaments[i].pemain[j].ACS)
			fmt.Println()
			p = tournaments[i].pemain[j]
			if p.ACS > MVP.ACS {
				MVP = p
			}
		}

	}

	fmt.Print("===Pemain Terbaik===\n")
	fmt.Printf("Pemain: %s\n", MVP.nama)
	fmt.Printf("Pemain: %d\n", MVP.ACS)

}
