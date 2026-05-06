package main

import "fmt"

type arr []int

func main() {
	var n int
	var array = make(arr, 100)
	var kelinci = array
	fmt.Print("masukan berapa jumlah kelinci:")
	fmt.Scan(&n)
	inputArr(n, &kelinci)
	var min, max int = minMax(kelinci, n)
	fmt.Println("berat kelinci terkecil:", min)
	fmt.Println("berat kelinci terbesar:", max)
}

func inputArr(n int, kelinci *arr) {
	var berat int
	for i := 0; i < n; i++ {
		fmt.Print("masukan berat kelinci:")
		fmt.Scan(&berat)
		(*kelinci)[i] = berat
	}
}

func minMax(kelinci arr, n int) (int, int) {
	var min, max int = kelinci[0], 0
	for _, vel := range kelinci[:n] {
		if min > vel {
			min = vel
		} else if max < vel {
			max = vel
		}
	}
	return min, max
}
