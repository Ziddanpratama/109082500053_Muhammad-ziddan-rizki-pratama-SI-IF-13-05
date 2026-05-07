package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cariNama string

	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")
	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cariNama)

	idxCepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat : %s\n", prov[idxCepat])

	fmt.Printf("Data provinsi yang dicari : %s\n", cariNama)

	Prediksi(prov, pop, tumbuh)
}
