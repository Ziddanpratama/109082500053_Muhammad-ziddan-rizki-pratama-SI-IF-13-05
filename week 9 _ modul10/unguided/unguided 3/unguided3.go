package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan berapa jumlah bayi:")
	fmt.Scan(&n)
	var array = make([]float32, n)
	var bayi = array
	inputArr(n, &bayi)
	var min, max float32 = minMax(bayi, n)
	fmt.Println("berat bayi terkecil:", min)
	fmt.Println("berat bayi terbesar:", max)
	fmt.Println("berat rata rata bayi:", average(bayi, n))
}

func inputArr(n int, bayi *[]float32) {
	var berat float32
	for i := 0; i < n; i++ {
		fmt.Print("masukan berat bayi:")
		fmt.Scan(&berat)
		(*bayi)[i] = berat
	}
}

func minMax(bayi []float32, n int) (float32, float32) {
	var min, max float32 = bayi[0], 0
	for _, vel := range bayi[:n] {
		if min > vel {
			min = vel
		} else if max < vel {
			max = vel
		}
	}
	return min, max
}

func average(bayi []float32, n int) float32 {
	var sum float32
	for _, vel := range bayi {
		sum += vel
	}
	return sum / float32(n)
}
