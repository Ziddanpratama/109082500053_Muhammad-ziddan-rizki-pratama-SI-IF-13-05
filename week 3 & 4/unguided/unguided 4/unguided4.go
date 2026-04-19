package main

import "fmt"

func hitungPersegi(s int) {
	fmt.Println("luas persegi :", s*s)
	fmt.Println("keliling persegi :", 4*s)

}
func hitungPersegiPanjang(p, l int) {
	fmt.Println("luas persegi panjang :", p*l)
	fmt.Println("keliling persegi panjang :", 2*(p+l))
}
func hitungLingkaran(r float64) {
	fmt.Println("luas lingkaran :", 3.14*r*r)
	fmt.Println("keliling lingkaran :", 2*3.14*r)
}
func main() {
	var pilih int
	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		var s int
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&s)
		hitungPersegi(s)
	case 2:
		var p, l int
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)
	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&r)
		hitungLingkaran(r)
	}
}
