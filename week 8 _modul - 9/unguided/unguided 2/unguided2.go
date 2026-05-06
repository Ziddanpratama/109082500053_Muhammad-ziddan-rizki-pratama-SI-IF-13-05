package main

import (
	"fmt"
	"math"
)

func masukanarray(n int, arr []int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen indeks ke-%d: ", i)
		fmt.Scan(&arr[i])
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen nilai: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	masukanarray(n, arr)

	fmt.Println("a. Keseluruhan isi array: ")
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()

	fmt.Println("b. Elemen dengan indeks ganjil: ")
	for _, vel := range arr {
		if vel%2 != 0 {
			fmt.Print(vel, " ")
		}
	}
	fmt.Println()

	fmt.Println("c. Elemen dengan indeks genap: ")
	for _, vel := range arr {
		if vel%2 == 0 {
			fmt.Print(vel, " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var deleteIdx int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&deleteIdx)
	arr = append(arr[:deleteIdx], arr[deleteIdx+1:]...)
	fmt.Print("e. Keseluruhan isi array setelah dihapus: ")
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()

	var sum float64
	for _, val := range arr {
		sum += float64(val)
	}
	var average float64
	if len(arr) > 0 {
		average = sum / float64(len(arr))
	}
	fmt.Printf("f. Rata-rata dari bilangan di dalam array: %.2f\n", average)

	var varianceSum float64
	for _, val := range arr {
		varianceSum += math.Pow(float64(val)-average, 2)
	}
	var stdDev float64
	if len(arr) > 0 {
		stdDev = math.Sqrt(varianceSum / float64(len(arr)))
	}
	fmt.Printf("g. Standar deviasi (simpangan baku): %.2f\n", stdDev)
}
