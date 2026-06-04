package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {

	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var input int

	fmt.Println("Masukan proses input suara :")
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}

		idx := posisi(p, n, input)
		if idx != -1 {
			p[idx].suara++
		} else {
			p[n].nama = input
			p[n].suara = 1
			n++
		}
	}

	for i := 1; i < n; i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < key.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	fmt.Println("Hasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
