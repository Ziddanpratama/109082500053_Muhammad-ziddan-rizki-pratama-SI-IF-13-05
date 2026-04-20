package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(faktorial(n))
}
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}
