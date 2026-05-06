package main

import (
	"fmt"
)

func main() {
	var A, B string
	var klubA, klubB []int
	fmt.Print("maskan klub A:")
	fmt.Scan(&A)
	fmt.Print("maskan klub B:")
	fmt.Scan(&B)
	var scoreA, scoreB int = 1, 1
	var i int = 1
	for scoreA > -1 && scoreB > -1 {
		fmt.Print("pertandingan ke ", i, ":")
		fmt.Scan(&scoreA, &scoreB)
		klubA = append(klubA, scoreA)
		klubB = append(klubB, scoreB)
		i++
	}
	var j int = 1
	for i := 0; i < len(klubA); i++ {
		if klubA[i] > klubB[i] {
			fmt.Println("hasil", j, ":", A)
		} else if klubA[i] < klubB[i] {
			fmt.Println("hasil", j, ":", B)
		} else {
			fmt.Println("hasil", j, "draw")
		}
		j++
	}
}
