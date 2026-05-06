package main

import "fmt"

func main() {
	var ikan [1000]float64
	var x, y int
	fmt.Print("Masukan Banyak Ikan Yang Dijual (x) : ")
	fmt.Scan(&x)
	fmt.Print("Masukan Kapasitas Ikan Per Wadah (y) : ")
	fmt.Scan(&y)
	var totalBerat float64 = 0
	for i := 0; i < x; i++ {
		fmt.Print("Masukan Berat Ikan Ke-", i+1, " : ")
		fmt.Scan(&ikan[i])
		totalBerat += ikan[i]
	}
	var beratWadah float64 = 0
	var jumlah float64 = 0
	fmt.Print("Total Berat Ikan di Setiap Wadah : ")

	for i := 0; i < x; i++ {
		beratWadah += ikan[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Print(beratWadah, " ")
			jumlah++
			beratWadah = 0
		}
	}

	fmt.Println()

	rataRata := totalBerat / jumlah
	fmt.Println("rata rata:", rataRata)
}
