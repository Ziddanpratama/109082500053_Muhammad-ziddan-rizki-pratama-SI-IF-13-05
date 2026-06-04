package main

import (
	"fmt"
)

const NMAX = 1000000 // Jumlah maksimum data masukan

type arrint [NMAX]int // Tipe data alias array integer

// SelectionSort mengurutkan array T secara membesar (ascending)
func SelectionSort(T *arrint, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

// median mengembalikan nilai median dari array T yang sudah terurut
func median(T arrint, n int) float64 {
	if n%2 == 1 {
		// Jika jumlah data ganjil, ambil nilai tengah
		return float64(T[n/2])
	} else {
		// Jika jumlah data genap, rata-rata dari dua nilai tengah
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
}

func main() {
	var A arrint
	var n int = 0
	var x int

	// Membaca data hingga bertemu -5313541
	for {
		fmt.Scan(&x)
		if x == -5313541 {
			break
		}

		if x == 0 {
			// Saat 0, urutkan data dan hitung median
			if n > 0 {
				// Kita perlu menyalin array karena selection sort mengubah urutan
				// Namun untuk efisiensi, kita bisa mengurutkan langsung di A
				temp := A // Salin data agar urutan asli tetap terjaga (opsional, tergantung kebutuhan)
				SelectionSort(&temp, n)
				fmt.Printf("%.1f\n", median(temp, n))
			}
		} else {
			// Masukkan data ke dalam array
			if n < NMAX {
				A[n] = x
				n++
			}
		}
	}
}
