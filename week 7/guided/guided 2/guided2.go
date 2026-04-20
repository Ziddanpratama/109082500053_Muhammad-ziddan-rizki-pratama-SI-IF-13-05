package main

import "fmt"

type Mahasiswa struct {
	Nama string
	NIM  string
	IPK  float64
}

func main() {
	var mhs [3]Mahasiswa

	for i := 0; i < 3; i++ {
		fmt.Println("Mahasiswa ke-", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&mhs[i].Nama)
		fmt.Print("NIM: ")
		fmt.Scan(&mhs[i].NIM)
		fmt.Print("IPK: ")
		fmt.Scan(&mhs[i].IPK)
	}
	fmt.Println("\nData Mahasiswa:")
	for i := 0; i < 3; i++ {
		fmt.Println(mhs[i])
	}
}
