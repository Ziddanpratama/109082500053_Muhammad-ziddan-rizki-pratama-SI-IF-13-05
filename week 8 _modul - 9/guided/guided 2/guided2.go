package main

import "fmt"

func main() {
	//deklarasi & inisialisasi map bernama nilai dengan key string dan value integer
	var nilai map[string]int = make(map[string]int)

	//Operasi CRUD (Create, Read, Update, Delete) & Searching

	//OPERASI CREATE
	nilai["Dhimas"] = 90 //Key = Dhimas, Value = 90
	nilai["Hafizh"] = 88 //Key = Hafizh, Value = 88
	nilai["Regita"] = 95 //Key = Regita, Value = 95

	//Operasi Read (membaca/menampilkan Data)
	fmt.Println("Data nilai : ")
	var nama string                 //variabel bantuan yang merujuk ke key
	var grade int                   //variabel bantuan yang merujuk ke value
	for nama, grade = range nilai { //untuk setiap nama (key) & grade (value) didalam range map nilai, maka tampilkan nama = nilai
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	//OPERASI UPDATE
	fmt.Println("Setelah Update : ")
	nilai["Regita"] = 92 //Update Value yang memiliki Key = Regita menjadi 92
	print("Regita = ", nilai["Regita"])
	fmt.Println()
	fmt.Println()

	//OPERASI DELETE
	delete(nilai, "Dhimas") //Hapus data didalam map nilai yang memiliki Key = Dhimas
	fmt.Println("Data nilai setelah delete : ")
	for nama, grade = range nilai {
		fmt.Println(nama, " = ", grade)
	}
	fmt.Println()

	//OPERASI SEARCHING
	fmt.Println("Hasil Searching : ")
	var cariNama string = "Hafizh" //variabel bantuan yang menyimpan data key (nama) yang akan dicari
	var isiValue int               //vaiabel bantuan yang merujuk ke value
	var ok bool                    //vaiabel bantuan untuk pengecekan true/false (boolean)
	isiValue, ok = nilai[cariNama] //ambil value yang memiliki Key = cariNama didalam map nilai, kemudian simpan kedalam isiValue & set ok menjadi True
	if ok {                        //jika ok bernilai True, maka tampilkan hasil searching nya (nilai cariNama = isiValue)
		fmt.Println("Nilai", cariNama, " = ", isiValue)
	} else { //jika ok bernilai false, maka tampilkan teks data tidak ditemukan
		fmt.Println("Data tidak ditemukan")
	}
}
