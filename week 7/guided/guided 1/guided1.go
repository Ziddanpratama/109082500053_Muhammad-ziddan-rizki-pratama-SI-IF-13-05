package main

import "fmt"

type Mahasiswa struct {
	Nama  string
	NIM   string
	IPK   float64
	Aktif bool
}

func main() {
	mhs1 := Mahasiswa{
		Nama:  "Budi Santoso",
		NIM:   "123456",
		IPK:   3.75,
		Aktif: true,
	}

	mhs2 := Mahasiswa{"Siti Rahayu", "789012", 3.90,
		true}
	var mhs3 Mahasiswa
	mhs3.Nama = "Andi Wijaya"
	mhs3.NIM = "345678"
	mhs3.IPK = 3.50
	mhs3.Aktif = false

	fmt.Println("Cara Inisialisasi 1:")
	fmt.Println(mhs1)
	fmt.Println("\nCara Inisialisasi 2:")
	fmt.Println(mhs2)
	fmt.Println("\nCara Inisialisasi 3:")
	fmt.Println(mhs3)
}
