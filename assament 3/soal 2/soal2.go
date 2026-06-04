package main

import (
	"fmt"
)

// Struct untuk menyimpan data pemain
type Pemain struct {
	Nama   string
	Gol    int
	Assist int
}

func main() {
	var n int
	fmt.Print("Masukan jumlah pemain: ")
	fmt.Scan(&n)

	pemain := make([]Pemain, n)

	// Input data pemain
	for i := 0; i < n; i++ {
		fmt.Scan(&pemain[i].Nama, &pemain[i].Gol, &pemain[i].Assist)
	}

	for i := 1; i < n; i++ {
		key := pemain[i]
		j := i - 1

		for j >= 0 && (pemain[j].Gol < key.Gol || (pemain[j].Gol == key.Gol && pemain[j].Assist < key.Assist)) {
			pemain[j+1] = pemain[j]
			j--
		}
		pemain[j+1] = key
	}

	fmt.Println("\nHasil Sorting :")
	for _, p := range pemain {
		fmt.Printf("%s %d %d\n", p.Nama, p.Gol, p.Assist)
	}
}
