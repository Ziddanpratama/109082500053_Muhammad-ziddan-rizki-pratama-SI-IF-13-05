package main

import "fmt"

func hitungTotalBelanja(harga int, jumlah int) {
	var total int
	total = harga * jumlah
	fmt.Println("total harga : ", total)

	hitungDiskon(total)
}
func hitungDiskon(total int) {
	var diskon, hargaAkhir int
	diskon = total * 10 / 100
	hargaAkhir = total - diskon

	fmt.Println("Diskon 10% : ", diskon)
	fmt.Println("harga akhir (setelah diskon): ",
		hargaAkhir)
}

func main() {
	var harga, jumlah int

	fmt.Print("masukkan harga : ")
	fmt.Scan(&harga)
	fmt.Print("masukkan jumlah : ")
	fmt.Scan(&jumlah)
	hitungTotalBelanja(harga, jumlah)
}
