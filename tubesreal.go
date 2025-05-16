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
		fmt.Println("5. Urutkan berdasarkan Selisih Skor (Insertion Sort)")
		fmt.Println("6. Cari Tim (Sequential Search)")
		fmt.Println("7. Cari Tim (Binary Search)")
		fmt.Println("8. Tampilkan Statistik")
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
			//sortByWins(&tournaments, totalTim) <- INI BISA DIGANTI VARIABLEnya, ALGORITMAnya, DLL-nya KALO MAU YAK! (shout out to case 4-8)
		case 5:
			//sortByScoreDiff(&tournaments, totalTim)
		case 6:
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&namaTeam)
			sequentialSearch(tournaments, totalTim, namaTeam)
		case 7:
			fmt.Print("Masukkan nama tim yang dicari: ")
			fmt.Scan(&namaTeam)
			binarySearch(tournaments, totalTim, namaTeam)
		case 8:
			//showStats(tournaments, totalTim)
		case 0:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("BUNYAMIN GOBLOK, YANG BENER LAH INPUTNYA WKWKWKWK")
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
		fmt.Print("Total SCA: ")
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

// EDIT TEAM
func editTeam(tournaments *[TMAX]team, totalTim int, namaTeam string) {
	var ketemu bool
	var i, j int
	var playerBaru, timBaru string
	var ACSUpdate int

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
			fmt.Scan(&tournaments[i].menang)
			fmt.Print("Selisih skor baru (0 jika tidak ingin mengubah): ")
			fmt.Scan(&tournaments[i].selisih)

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
	i = 0
	for i < *totalTim && !found {
		if tournaments[i].nama == teamName {
			for j = i; j < *totalTim-1; j++ {
				tournaments[j] = tournaments[j+1]
			}
			*totalTim--
			fmt.Println("Tim berhasil dihapus!")
			found = true
		}
		i++
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
		fmt.Println("\nTim dengan nama", namaTeam, "gk ada beb")
	}
}

// BINARY SEARCH CARI NAMA TIM
func binarySearch(tournaments [TMAX]team, totalTim int, namaTeam string) {
	var i int
	var left, right, mid, hasil int
	var ketemu bool

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
		fmt.Println("\nTim dengan nama", namaTeam, "gk ada beb")
	}
}
