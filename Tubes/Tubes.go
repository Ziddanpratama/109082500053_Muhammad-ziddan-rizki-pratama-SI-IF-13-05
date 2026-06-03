package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	ID   int
	Name string
}

type Medicine struct {
	ID            int
	Name          string
	CategoryID    int
	Indication    string
	Stock         int
	ExpiryDate    time.Time
	IncomingDate  time.Time
	Unit          string
	MinStockAlert int
}

var reader = bufio.NewReader(os.Stdin)

var categories = []Category{
	{ID: 1, Name: "Demam"},
	{ID: 2, Name: "Batuk"},
	{ID: 3, Name: "Nyeri"},
}

var medicines = []Medicine{}
var nextMedicineID = 1
var nextCategoryID = 4

func main() {
	for {
		fmt.Println("\n==============================")
		fmt.Println("Aplikasi Apotek-Smart")
		fmt.Println("==============================")
		fmt.Println("1. Kelola Kategori Gejala")
		fmt.Println("2. Kelola Data Obat")
		fmt.Println("3. Tambah Transaksi Stok Masuk")
		fmt.Println("4. Cari Obat")
		fmt.Println("5. Urutkan Data Obat")
		fmt.Println("6. Statistik Stok")
		fmt.Println("7. Tampilkan Semua Data Obat")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		menu := readInt()
		switch menu {
		case 1:
			categoryMenu()
		case 2:
			medicineMenu()
		case 3:
			addStockTransaction()
		case 4:
			searchMenu()
		case 5:
			sortMenu()
		case 6:
			showStatistics()
		case 7:
			listMedicines()
		case 0:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func categoryMenu() {
	for {
		fmt.Println("\n--- Kelola Kategori ---")
		fmt.Println("1. Tambah kategori")
		fmt.Println("2. Ubah kategori")
		fmt.Println("3. Hapus kategori")
		fmt.Println("4. Lihat kategori")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		choice := readInt()
		switch choice {
		case 1:
			fmt.Print("Nama kategori: ")
			name := readString()
			categories = append(categories, Category{ID: nextCategoryID, Name: name})
			nextCategoryID++
			fmt.Println("Kategori berhasil ditambahkan.")
		case 2:
			showCategories()
			fmt.Print("ID kategori yang diubah: ")
			id := readInt()
			for i := range categories {
				if categories[i].ID == id {
					fmt.Print("Nama baru: ")
					categories[i].Name = readString()
					fmt.Println("Kategori berhasil diubah.")
					goto endCategoryMenu
				}
			}
			fmt.Println("Kategori tidak ditemukan.")
		case 3:
			showCategories()
			fmt.Print("ID kategori yang dihapus: ")
			id := readInt()
			for i := range categories {
				if categories[i].ID == id {
					categories = append(categories[:i], categories[i+1:]...)
					fmt.Println("Kategori berhasil dihapus.")
					goto endCategoryMenu
				}
			}
			fmt.Println("Kategori tidak ditemukan.")
		case 4:
			showCategories()
		case 0:
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	endCategoryMenu:
	}
}

func medicineMenu() {
	for {
		fmt.Println("\n--- Kelola Data Obat ---")
		fmt.Println("1. Tambah obat")
		fmt.Println("2. Ubah obat")
		fmt.Println("3. Hapus obat")
		fmt.Println("4. Lihat obat")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		choice := readInt()
		switch choice {
		case 1:
			addMedicine()
		case 2:
			updateMedicine()
		case 3:
			deleteMedicine()
		case 4:
			listMedicines()
		case 0:
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func addMedicine() {
	showCategories()
	fmt.Print("Nama obat: ")
	name := readString()

	fmt.Print("ID kategori: ")
	categoryID := readInt()

	if !categoryExists(categoryID) {
		fmt.Println("Kategori tidak valid.")
		return
	}

	fmt.Print("Indikasi/gejala: ")
	indication := readString()

	fmt.Print("Stok: ")
	stock := readInt()

	fmt.Print("Satuan: ")
	unit := readString()

	fmt.Print("Tanggal kedaluwarsa (YYYY-MM-DD): ")
	expiry := readDate()

	fmt.Print("Tanggal masuk stok (YYYY-MM-DD): ")
	incoming := readDate()

	fmt.Print("Batas stok minimum: ")
	minStock := readInt()

	med := Medicine{
		ID:            nextMedicineID,
		Name:          name,
		CategoryID:    categoryID,
		Indication:    indication,
		Stock:         stock,
		ExpiryDate:    expiry,
		IncomingDate:  incoming,
		Unit:          unit,
		MinStockAlert: minStock,
	}
	medicines = append(medicines, med)
	nextMedicineID++
	fmt.Println("Obat berhasil ditambahkan.")
}

func updateMedicine() {
	listMedicines()
	fmt.Print("ID obat yang diubah: ")
	id := readInt()

	for i := range medicines {
		if medicines[i].ID == id {
			fmt.Print("Nama obat baru: ")
			medicines[i].Name = readString()

			showCategories()
			fmt.Print("ID kategori baru: ")
			cid := readInt()
			if categoryExists(cid) {
				medicines[i].CategoryID = cid
			}

			fmt.Print("Indikasi baru: ")
			medicines[i].Indication = readString()

			fmt.Print("Stok baru: ")
			medicines[i].Stock = readInt()

			fmt.Print("Satuan baru: ")
			medicines[i].Unit = readString()

			fmt.Print("Tanggal kedaluwarsa baru (YYYY-MM-DD): ")
			medicines[i].ExpiryDate = readDate()

			fmt.Print("Tanggal masuk baru (YYYY-MM-DD): ")
			medicines[i].IncomingDate = readDate()

			fmt.Print("Batas stok minimum baru: ")
			medicines[i].MinStockAlert = readInt()

			fmt.Println("Obat berhasil diubah.")
			return
		}
	}
	fmt.Println("Obat tidak ditemukan.")
}

func deleteMedicine() {
	listMedicines()
	fmt.Print("ID obat yang dihapus: ")
	id := readInt()

	for i := range medicines {
		if medicines[i].ID == id {
			medicines = append(medicines[:i], medicines[i+1:]...)
			fmt.Println("Obat berhasil dihapus.")
			return
		}
	}
	fmt.Println("Obat tidak ditemukan.")
}

func addStockTransaction() {
	listMedicines()
	fmt.Print("ID obat: ")
	id := readInt()

	for i := range medicines {
		if medicines[i].ID == id {
			fmt.Print("Jumlah stok masuk: ")
			qty := readInt()
			medicines[i].Stock += qty
			fmt.Print("Tanggal transaksi (YYYY-MM-DD): ")
			tgl := readDate()
			fmt.Printf("Transaksi stok masuk berhasil pada %s.\n", tgl.Format("2006-01-02"))
			return
		}
	}
	fmt.Println("Obat tidak ditemukan.")
}

func searchMenu() {
	for {
		fmt.Println("\n--- Cari Obat ---")
		fmt.Println("1. Sequential Search berdasarkan nama")
		fmt.Println("2. Sequential Search berdasarkan indikasi")
		fmt.Println("3. Binary Search berdasarkan nama")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		choice := readInt()
		switch choice {
		case 1:
			fmt.Print("Masukkan nama obat: ")
			key := strings.ToLower(readString())
			results := sequentialSearchByName(key)
			printMedicineResults(results)
		case 2:
			fmt.Print("Masukkan indikasi: ")
			key := strings.ToLower(readString())
			results := sequentialSearchByIndication(key)
			printMedicineResults(results)
		case 3:
			fmt.Print("Masukkan nama obat: ")
			key := strings.ToLower(readString())
			result := binarySearchByName(key)
			if result == nil {
				fmt.Println("Obat tidak ditemukan.")
			} else {
				printMedicine(*result)
			}
		case 0:
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func sortMenu() {
	for {
		fmt.Println("\n--- Urutkan Data Obat ---")
		fmt.Println("1. Selection Sort berdasarkan tanggal kedaluwarsa")
		fmt.Println("2. Insertion Sort berdasarkan tanggal kedaluwarsa")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		choice := readInt()
		switch choice {
		case 1:
			sorted := selectionSortByExpiry(copyMedicines())
			printMedicineResults(sorted)
		case 2:
			sorted := insertionSortByExpiry(copyMedicines())
			printMedicineResults(sorted)
		case 0:
			return
		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}

func showStatistics() {
	nearEmpty := 0
	expiringSoon := []Medicine{}
	now := time.Now()
	limit := now.AddDate(0, 1, 0)

	for _, m := range medicines {
		if m.Stock <= m.MinStockAlert {
			nearEmpty++
		}
		if !m.ExpiryDate.After(limit) {
			expiringSoon = append(expiringSoon, m)
		}
	}

	fmt.Println("\n--- Statistik Stok ---")
	fmt.Printf("Jumlah obat hampir habis: %d\n", nearEmpty)
	fmt.Println("Daftar obat yang segera kedaluwarsa:")
	if len(expiringSoon) == 0 {
		fmt.Println("- Tidak ada")
	} else {
		for _, m := range expiringSoon {
			fmt.Printf("- %s | Exp: %s | Stok: %d\n", m.Name, m.ExpiryDate.Format("2006-01-02"), m.Stock)
		}
	}
}

func listMedicines() {
	if len(medicines) == 0 {
		fmt.Println("Data obat masih kosong.")
		return
	}
	fmt.Println("\n--- Daftar Obat ---")
	for _, m := range medicines {
		printMedicine(m)
	}
}

func printMedicine(m Medicine) {
	catName := getCategoryName(m.CategoryID)
	fmt.Printf("ID: %d | Nama: %s | Kategori: %s | Indikasi: %s | Stok: %d %s | Exp: %s | Masuk: %s | Min: %d\n",
		m.ID, m.Name, catName, m.Indication, m.Stock, m.Unit,
		m.ExpiryDate.Format("2006-01-02"),
		m.IncomingDate.Format("2006-01-02"),
		m.MinStockAlert)
}

func printMedicineResults(results []Medicine) {
	if len(results) == 0 {
		fmt.Println("Tidak ada data yang ditemukan.")
		return
	}
	for _, m := range results {
		printMedicine(m)
	}
}

func sequentialSearchByName(key string) []Medicine {
	var results []Medicine
	for _, m := range medicines {
		if strings.Contains(strings.ToLower(m.Name), key) {
			results = append(results, m)
		}
	}
	return results
}

func sequentialSearchByIndication(key string) []Medicine {
	var results []Medicine
	for _, m := range medicines {
		if strings.Contains(strings.ToLower(m.Indication), key) {
			results = append(results, m)
		}
	}
	return results
}

func binarySearchByName(key string) *Medicine {
	sorted := copyMedicines()
	sort.Slice(sorted, func(i, j int) bool {
		return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
	})

	low, high := 0, len(sorted)-1
	for low <= high {
		mid := (low + high) / 2
		name := strings.ToLower(sorted[mid].Name)
		if name == key {
			return &sorted[mid]
		} else if name < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func selectionSortByExpiry(data []Medicine) []Medicine {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].ExpiryDate.Before(data[minIdx].ExpiryDate) {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	return data
}

func insertionSortByExpiry(data []Medicine) []Medicine {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].ExpiryDate.After(key.ExpiryDate) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	return data
}

func copyMedicines() []Medicine {
	cp := make([]Medicine, len(medicines))
	copy(cp, medicines)
	return cp
}

func showCategories() {
	fmt.Println("\n--- Kategori ---")
	for _, c := range categories {
		fmt.Printf("ID: %d | Nama: %s\n", c.ID, c.Name)
	}
}

func categoryExists(id int) bool {
	for _, c := range categories {
		if c.ID == id {
			return true
		}
	}
	return false
}

func getCategoryName(id int) string {
	for _, c := range categories {
		if c.ID == id {
			return c.Name
		}
	}
	return "Tidak diketahui"
}

func readString() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt() int {
	for {
		text := readString()
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Print("Input angka valid: ")
	}
}

func readDate() time.Time {
	for {
		text := readString()
		t, err := time.Parse("2006-01-02", text)
		if err == nil {
			return t
		}
		fmt.Print("Format salah. Masukkan YYYY-MM-DD: ")
	}
}
