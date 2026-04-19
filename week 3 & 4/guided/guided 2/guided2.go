package main

import "fmt"

func tambahValue(x int) {
	x = x + 10
	fmt.Println("Nilai x didalam prosdur tambahValue (pass by value) : ", x)
}

func tambahReference(x *int) {
	*x = *x + 10
	fmt.Println("Nilai x didalam prosdur tambahReference (pass by refeence) : ", *x)
}

func main() {
	var y int = 5

	fmt.Println("nilai y awal : ", y)
	tambahValue(y)
	fmt.Println("nilai y setelah pass by value : ", y)
	tambahReference(&y)
	fmt.Println("nilai y setelah pass by reference : ", y)
}
