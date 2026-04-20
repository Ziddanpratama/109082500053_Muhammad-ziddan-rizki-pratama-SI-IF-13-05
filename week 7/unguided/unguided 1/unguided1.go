package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return (4.0 / 5.0) * c
}
func CelciusToFahrenheit(c suhu) suhu {
	return (9.0/5.0)*c + 32
}
func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}
func main() {
	var c suhu
	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&c)
	r := CelciusToReamur(c)
	f := CelciusToFahrenheit(c)
	k := CelciusToKelvin(c)

	fmt.Println()
	fmt.Println(c, "celcius =", r, "reamur")
	fmt.Println(c, "celcius =", f, "fahrenheit")
	fmt.Println(c, "celcius =", k, "kelvin")
}
