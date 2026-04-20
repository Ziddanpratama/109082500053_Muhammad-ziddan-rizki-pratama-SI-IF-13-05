package main

import "fmt"

type angka int
type kata string
type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var b Buku
	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukkan judul buku :")
	fmt.Scan(&b.judul)
	fmt.Print("Masukkan nama penulis :")
	fmt.Scan(&b.penulis)
	fmt.Print("Masukkan penerbit : ")
	fmt.Scan(&b.penerbit)
	fmt.Print("Masukkan tahun terbit :")
	fmt.Scan(&b.tahunTerbit)
	fmt.Print("Masukkan jumlah halaman :")
	fmt.Scan(&b.jumlahHalaman)
	fmt.Println()
	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", b.judul)
	fmt.Println("Penulis :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun Terbit :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}
