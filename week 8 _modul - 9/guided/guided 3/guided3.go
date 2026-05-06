package main

import "fmt"

func main() {
	var nilai []int //slice = dinamis (ukurannya tidak ditentukan)

	//OPERASI CRUD (Create, Read, Update, Delete) & Searching

	//OPERASI CREATE MANUAL PAKAI APPEND

	//OPERASI CREATE PAKAI FOR
	var jmlData int = 3
	var idx int
	for idx = 0; idx < jmlData; idx++ {
		var inputNilai int
		fmt.Printf("Masukkan nilai ke-%d : ", idx)
		fmt.Scan(&inputNilai)
		nilai = append(nilai, inputNilai)
	}
	fmt.Println()

	//OPERASI READ
	fmt.Println("Semua data:", nilai)
	fmt.Println("Slice [0:2]:", nilai[:2]) //print data didalam slice nilai dari index 0 sampai sebelum index 2
	fmt.Println("Panjang slice:", len(nilai))

	//OPERASI UPDATE
	nilai[0] = 100
	fmt.Println("\nSetelah update:", nilai)

	//OPERASI DELETE
	var indexHapus int = 1
	nilai = append(nilai[:indexHapus], nilai[indexHapus+1:]...) //ambil nilai sebelum indexHapus dan nilai setelah indexHapus dan gabungkan kedalam slice nilai

	fmt.Println("\nSetelah delete:", nilai)
	fmt.Println()

	//OPERASI SEARCHING
	var cariNilai int = 75
	var found bool = false
	var i int

	for i = 0; i < len(nilai); i++ {
		if nilai[i] == cariNilai {
			fmt.Printf("Nilai %d ditemukan di index %d\n", cariNilai, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Nilai tidak ditemukan")
	}
}
