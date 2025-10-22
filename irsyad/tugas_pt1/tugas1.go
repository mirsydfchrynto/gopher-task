package main

import "fmt"

func main() {
	
	kalkulator("perkalian", 13, 2)
	kalkulator("pembagian", 13, 2)
	kalkulator("penjumlahan", 13, 2)
	kalkulator("pengurangan", 13, 2)
	kalkulator("testDefault", 13,18)
}

func kalkulator(operasi string, angka1, angka2 int) {
	var hasilnya int
	switch operasi {
	case "penjumlahan":
		hasilnya = angka1 + angka2
		fmt.Printf("operasi pnjumlahan dari %v + %v hasilnya adalah %v\n", angka1, angka2, hasilnya)
	case "pengurangan":
		hasilnya = angka1 - angka2
		fmt.Printf("operasi pengurangan dari %v - %v hasilnya adalah %v\n", angka1, angka2, hasilnya)
	case "pembagian":
		hasilnya = angka1 / angka2
		fmt.Printf("operasi pembagian dari %v : %v hasilnya adalah %v\n", angka1, angka2, hasilnya)
	case "perkalian":
		hasilnya = angka1 * angka2
		fmt.Printf("operasi perkalian dari %v x %v hasilnya adalah %v\n", angka1, angka2, hasilnya)
	default:
		fmt.Println("tidak ada operasi perhitungan")
	}
}
