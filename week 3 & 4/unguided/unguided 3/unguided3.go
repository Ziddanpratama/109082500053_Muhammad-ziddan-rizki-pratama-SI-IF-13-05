package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n, " ")
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}
func main() {
	var n int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	cetakDeret(n)
}
