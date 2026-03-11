package main

import "fmt"

func main() {
	var pilihan, angka int

	fmt.Println("-- MENU --")
	fmt.Println("1. cek angka == 10")
	fmt.Println("2. cek genap ganjil")
	fmt.Print("masukkan pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print("masukkan angka : ")
		fmt.Scan(&angka)
		if angka == 10 {
			fmt.Println("angka adalah 10")
		} else {
			fmt.Println("angka bukan 10")
		}
	case 2:
		fmt.Print("masukkan angka : ")
		fmt.Scan(&angka)
		if angka%2 == 0 {
			fmt.Println("angka genap")
		} else if angka%2 != 0 {
			fmt.Println("angka ganjil")
		}
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
