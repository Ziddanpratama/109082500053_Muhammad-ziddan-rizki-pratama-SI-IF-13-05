package main

import "fmt"

func main() {
	var mahasiswa [3]string //bisa menyimpan 3 data

	//OPERASI CRUD (Create, Read, Update, Delete) & Searching

	//CREATE MANUAL
	// mahasiswa[0] = "Dhimas"
	// mahasiswa[1] = "Hafizh"

	//CREATE PAKAI FOR
	var i int
	for i = 0; i < 3; i++ {
		fmt.Printf("Masukkan data index ke-%d : ", i)
		fmt.Scan(&mahasiswa[i])
	}

	//OPERASI READ MANUAL
	// fmt.Println("index ke-0 : ", mahasiswa[0])
	// fmt.Println("index ke-1 : ", mahasiswa[1])

	//OPERASI READ PAKAI FOR
	var j int
	for j = 0; j < 3; j++ {
		fmt.Println("data ke-", j, " : ", mahasiswa[j])
	}

	//READ PAKAI SLICE
	fmt.Println("index [:3] : ", mahasiswa[:3]) //print isi array dari index 0 sampai index sebelum 3 (0, 1, 2)
	fmt.Println("index [:1] : ", mahasiswa[1:]) //print isi array dari index ke 1 sampai akhir

	fmt.Println(len(mahasiswa)) //menampilkan ukuran array

	//OPERASI UPDATE
	mahasiswa[1] = "Regita"
	fmt.Println(mahasiswa[1])

	//OPERASI DELETE
	var indexHapus = 0
	var idx int
	for idx = indexHapus; idx < len(mahasiswa)-1; idx++ {
		mahasiswa[idx] = mahasiswa[idx+1]
	}
	mahasiswa[len(mahasiswa)-1] = ""

	fmt.Println(mahasiswa[:3])

	//OPERASI SEARCHING
	var cariNama string = "Regita"
	var found bool = false
	var k int
	for k = 0; k < 3; k++ {
		if mahasiswa[k] == cariNama {
			fmt.Printf("data ditemukan pada index ke-%d", k)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("data tidak ditemukan")
	}
}
