package main

import "fmt"

type arrInt [2023]int

func terkecil_1(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var j int = 1

	for j < n {
		if min > tabInt[j] {
			min = tabInt[j]
		}
		j = j + 1
	}
	return min
}

func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0
	var j int = 1

	for j < n {
		if tabInt[idx] > tabInt[j] { // [9, 8, 5, 6]
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var arr arrInt

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen array :")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	min := terkecil_1(arr, n)
	idx := terkecil_2(arr, n)

	fmt.Println("Nilai terkecil: ", min)
	fmt.Println("Indeks nilai terkecil: ", idx)
}
