package main

import "fmt"

type mahasiswa struct {
	nama string
	nim  string
	IPK  float64
}

type arrMahasiswa [3]mahasiswa

//data index 0 = nama, nim, IPK
//data index 1 = nama, nim, ipk
//data index 2 = nama, nim, ipk

func IPKTerbesar(array []mahasiswa) (mahasiswa, int) {
	var mhsTerbesar mahasiswa = array[0] //nilai awal
	var idxDitemukan int = 0
	var i int

	for i = 1; i < len(array); i++ {
		if array[i].IPK > mhsTerbesar.IPK {
			mhsTerbesar = array[i]
			idxDitemukan = i
		}
	}

	return mhsTerbesar, idxDitemukan
}

func IPKTerkecil(array []mahasiswa) (mahasiswa, int) {
	var mhsTerkecil mahasiswa = array[0] //nilai awal
	var idxDitemukan int = 0
	var i int

	for i = 1; i < len(array); i++ {
		if array[i].IPK < mhsTerkecil.IPK {
			mhsTerkecil = array[i]
			idxDitemukan = i
		}
	}

	return mhsTerkecil, idxDitemukan
}

func main() {
	var arrMhs arrMahasiswa
	var i int

	for i = 0; i < len(arrMhs); i++ {
		fmt.Println("INPUT DATA MAHASISWA INDEX KE-", i)
		fmt.Print("Masukkan nama : ")
		fmt.Scan(&arrMhs[i].nama)
		fmt.Print("Masukkkan nim : ")
		fmt.Scan(&arrMhs[i].nim)
		fmt.Print("Masukkan IPK : ")
		fmt.Scan(&arrMhs[i].IPK)
		fmt.Println("==================================")
	}
	fmt.Println()

	var mhsMaks, mhsMin mahasiswa
	var idxMaks, idxMin int

	mhsMaks, idxMaks = IPKTerbesar(arrMhs[:])

	mhsMin, idxMin = IPKTerkecil(arrMhs[:])

	fmt.Println("DATA IPK TERKECIL DAN TERBESAR")
	fmt.Println("=== IPK TERBESAR ===")
	fmt.Println("IPK Terbesar = ", mhsMaks.IPK)
	fmt.Println("Atas nama = ", mhsMaks.nama)
	fmt.Println("nim = ", mhsMaks.nim)
	fmt.Println("ditemukan pada indeks ke-", idxMaks)
	fmt.Println()
	fmt.Println("=== IPK TERKECIL ===")
	fmt.Println("IPK Terkecil = ", mhsMin.IPK)
	fmt.Println("Atas nama = ", mhsMin.nama)
	fmt.Println("nim = ", mhsMin.nim)
	fmt.Println("ditemukan pada indeks ke-", idxMin)
}
