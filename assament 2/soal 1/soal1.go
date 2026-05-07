package main

import "fmt"

type setT [2022]int

func exist(arr setT, n int, val int) bool {
	for i := 0; i < n; i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func inputSet(arr *setT, n int) int {
	var x int
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		if exist(*arr, n, x) {
			break
		}
		arr[n] = x
		n++
	}
	return n
}

func findIntersection(T1, T2 setT, n1, n2 int, T3 *setT, n3 *int) {
	*n3 = 0
	for i := 0; i < n1; i++ {
		if exist(T2, n2, T1[i]) {
			T3[*n3] = T1[i]
			(*n3)++
		}
	}
}

func printSet(T setT, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 setT
	var n1, n2, n3 int
	n1 = inputSet(&s1, 0)
	n2 = inputSet(&s2, 0)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
