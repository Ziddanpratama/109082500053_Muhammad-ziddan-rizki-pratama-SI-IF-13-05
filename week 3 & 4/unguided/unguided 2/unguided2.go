package main

import "fmt"

func hitungBiaya(jenis string, masuk, keluar int) int {
	var total int = 0
	var jam int
	for jam = masuk; jam < keluar; jam++ {
		if jenis == "motor" {
			if jam < 17 {
				total = total + 4000
			} else {
				total = total + 5000
			}
		} else if jenis == "mobil" {
			if jam < 17 {
				total = total + 6000
			} else {
				total = total + 7000
			}
		}
	}
	return total
}

func main() {
	var jenis string
	var masuk, keluar int
	var total int = 0
	var no int = 1

	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====")

	for {
		fmt.Println("\n*Kendaraan", no)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)

		fmt.Print("Masukkan Jam Keluar Kendaraan (024): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)

		fmt.Println("Biaya parkir", jenis, no, ":", biaya)
		fmt.Println("==========================================")

		total = total + biaya
		no = no + 1
	}

	fmt.Println("\n*** Total Pendapatan Parkir Hari Ini Adalah", total, "***")
}
    