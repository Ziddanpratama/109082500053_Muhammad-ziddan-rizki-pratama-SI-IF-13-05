package main

import "fmt"

func main() {
	var angka int
	fmt.Print("masukkan angka : ")
	fmt.Scan(&angka)

	for angka < 10 {
		angka++
		fmt.Println("angka sekarang", angka)
	}
}
