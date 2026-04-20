package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n, 1)
}
func bintang(n int, i int) {
	if i > n {
		return
	} else {
		cetakBintang(i)
		fmt.Println()
		bintang(n, i+1)
	}
}
func cetakBintang(jumlah int) {
	if jumlah == 0 {
		return
	} else {
		fmt.Print("*")
		cetakBintang(jumlah - 1)
	}
}
