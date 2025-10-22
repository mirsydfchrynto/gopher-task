package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	jalankan()
}

func input(ketik string, tujuan *string) {
	fmt.Print(ketik)
	inputan, _ := reader.ReadString('\n')
	*tujuan = strings.TrimSpace(inputan)
}

type Data struct {
	judul  string
	isi    string
	status string
}

type Database struct {
	data []Data
}


func tambahdata(db *Database) {
	var tambah Data
	input("Masukan judul data: ", &tambah.judul)
	input("Masukan isi data: ", &tambah.isi)
	tambah.status = "belum selesai"
	db.data = append(db.data, tambah)
	fmt.Println("Data berhasil ditambahkan.")
	fmt.Println()
}

func statusdata(db *Database) {
	if len(db.data) == 0 {
		fmt.Println("Belum ada data untuk diubah.")
		fmt.Println()
	} else {
		lihatdata(db)
		var pilih string
		input("Masukan judul data yang ingin diubah statusnya: ", &pilih)

		ketemu := false
		for i := 0; i < len(db.data); i++ {
			if db.data[i].judul == pilih {
				if db.data[i].status == "selesai" {
					db.data[i].status = "belum selesai"
				} else {
					db.data[i].status = "selesai"
				}
				fmt.Println("Status data berhasil diubah.")
				ketemu = true
				break
			}
		}

		if !ketemu {
			fmt.Println("Data dengan judul tersebut tidak ditemukan.")
		}
		fmt.Println()
	}
}

func lihatdata(db *Database) {
	fmt.Println("=== Daftar Data ===")
	if len(db.data) == 0 {
		fmt.Println("Belum ada data.")
	} else {
		for i := 0; i < len(db.data); i++ {
			fmt.Println()
			fmt.Println("Judul :", db.data[i].judul)
			fmt.Println("Isi   :", db.data[i].isi)
			fmt.Println("Status:", db.data[i].status)
			fmt.Println("-------------------")
		}
	}
	fmt.Println()
}

func hapusdata(db *Database) {
	if len(db.data) == 0 {
		fmt.Println("Belum ada data untuk dihapus.")
		fmt.Println()
	} else {
		var pilih string
		input("Masukan judul data yang ingin dihapus: ", &pilih)

		ketemu := false
		for i := 0; i < len(db.data); i++ {
			if db.data[i].judul == pilih {
				db.data = append(db.data[:i], db.data[i+1:]...)
				fmt.Println("Data berhasil dihapus.")
				ketemu = true
				break
			}
		}

		if !ketemu {
			fmt.Println("Data dengan judul tersebut tidak ditemukan.")
		}
		fmt.Println()
	}
}

func updatedata(db *Database) {
	for i := 0; i < len(db.data); i++ {
			fmt.Println("-", db.data[i].judul)
	}
	if len(db.data) == 0 {
		fmt.Println("Belum ada data untuk diupdate.")
		fmt.Println()
	} else {
		var pilih string
		input("Masukan judul data yang ingin diupdate: ", &pilih)

		ketemu := false
		for i := 0; i < len(db.data); i++ {
			if db.data[i].judul == pilih {
				input("Masukan judul baru: ", &db.data[i].judul)
				input("Masukan isi baru: ", &db.data[i].isi)
				fmt.Println("Data berhasil diupdate.")
				ketemu = true
				break
			}
		}

		if !ketemu {
			fmt.Println("Data dengan judul tersebut tidak ditemukan.")
		}
		fmt.Println()
	}
}

func menu() {
	fmt.Println("=== TODO CLI APP ===")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Lihat Data")
	fmt.Println("3. Update Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Ubah Status Data")
	fmt.Println("6. Keluar")
}

func jalankan() {
	var db Database
	var pilihan string

	for {
		menu()
		input("Pilih menu (1-6): ", &pilihan)
		fmt.Println()

		if pilihan == "1" {
			tambahdata(&db)
		} else if pilihan == "2" {
			lihatdata(&db)
		} else if pilihan == "3" {
			updatedata(&db)
		} else if pilihan == "4" {
			hapusdata(&db)
		} else if pilihan == "5" {
			statusdata(&db)
		} else if pilihan == "6" {
			break
		} else {
			fmt.Println("pilihan tidak valid, silakan coba lagi.")
			fmt.Println()
		}
	}
}
