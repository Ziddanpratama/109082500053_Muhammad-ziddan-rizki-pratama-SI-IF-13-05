package main

import (
	"fmt"
)

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func main() {
	var T arrayMahasiswa
	var n int
	var searchNIM string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&searchNIM)

	idxPertama := cariNilaiPertama(T, n, searchNIM)
	if idxPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", searchNIM, T[idxPertama].nilai)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", searchNIM, cariNilaiTerbesar(T, n, searchNIM))
	} else {
		fmt.Println("Data mahasiswa dengan NIM tersebut tidak ditemukan.")
	}
}
func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			return i
		}
	}
	return -1
}
func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > max {
				max = T[i].nilai
			}
		}
	}
	return max
}
