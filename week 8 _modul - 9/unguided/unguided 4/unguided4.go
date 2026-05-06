package main

import "fmt"

const max int = 127

type tabel [max]rune

func main() {
	var m int
	var tab tabel
	input(&tab, &m)
	t := balikanArr(&tab, &m)
	cetakArr(t)
	fmt.Println(palidrom(tab, t))
}

func input(t *tabel, n *int) {
	var input string
	for *n < max {
		fmt.Scan(&input)
		char := rune(input[0])
		if char == '.' {
			break
		}
		t[*n] = char
		(*n)++
	}
}

func cetakArr(t tabel) {
	for _, vel := range t {
		fmt.Printf("%c", vel)
	}
	fmt.Println()
}

func balikanArr(t *tabel, n *int) tabel {
	var balik tabel
	var j int
	for i := *n - 1; i >= 0; i-- {
		char := (*t)[i]
		balik[j] = char
		j++
	}
	return balik
}

func palidrom(t tabel, T tabel) bool {
	if t == T {
		return true
	} else {
		return false
	}
}
