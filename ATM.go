package main

import (
	"fmt"
)

var Username string = "Helga"
var Password string = "2406499102"
var saldo int = 3500000
var historitransaksi []string

var inputUsername, inputPassword string

func main() {
	fmt.Print("Masukkan username: ")
	fmt.Scan(&inputUsername)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&inputPassword)

	if inputUsername == Username && inputPassword == Password {
		pilihanMenu := []string{
			"Lihat Informasi Akun",
			"Lihat Saldo",
			"Tambahkan Saldo",
			"Tarik Saldo",
			"Histori Transaksi",
			"Keluar dari Program",
		}

		var pilihan int
		for {
			displayMenu(pilihanMenu)
			fmt.Print("Masukkan pilihan menu: ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				LihatInformasiAkun()
			case 2:
				lihatSaldo()
			case 3:
				tambahSaldo()
			case 4:
				tarikSaldo()
			case 5:
				lihatHistoritransaksi()
			case 6:
				fmt.Println("Terima kasih telah menggunakan ATM!")
				return
			default:
				fmt.Println("Pilihan tidak valid, coba lagi.")
			}
		}
	} else {
		fmt.Println("Username atau password salah.")
	}
}

func LihatInformasiAkun() {
	fmt.Println("\n--- Informasi Akun ---")
	fmt.Printf("Username: %s\n", Username)
	fmt.Printf("Saldo Terkini: Rp %d\n", saldo)
}

func lihatSaldo() {
	fmt.Printf("Saldo kamu saat ini: Rp %d\n", saldo)
}

func tambahSaldo() {
	var jumlah int
	fmt.Print("Masukkan jumlah saldo yang ingin ditambahkan: ")
	fmt.Scan(&jumlah)
	if jumlah <= 0 {
		fmt.Println("Penambahan saldo harus lebih besar dari 0.")
	} else {
		saldo += jumlah
		historitransaksi = append(historitransaksi, fmt.Sprintf("Penambahan: Rp %d", jumlah))
		fmt.Printf("Saldo berhasil ditambahkan.\n")
	}
}

func tarikSaldo() {
	var jumlah int
	fmt.Print("Masukkan jumlah saldo yang ingin ditarik: ")
	fmt.Scan(&jumlah)

	if jumlah > saldo {
		fmt.Println("Saldo tidak cukup!")
	} else if jumlah <= 0 {
		fmt.Println("Jumlah penarikan harus lebih besar dari 0.")
	} else {
		saldo -= jumlah
		historitransaksi = append(historitransaksi, fmt.Sprintf("Penarikan: Rp %d", jumlah))
		fmt.Printf("Saldo berhasil ditarik. Saldo baru: Rp %d\n", saldo)
	}
}

func lihatHistoritransaksi() {
	fmt.Println("\n--- Histori Transaksi ---")
	if len(historitransaksi) == 0 {
		fmt.Println("Belum ada transaksi.")
	} else {
		for _, transaksi := range historitransaksi {
			fmt.Println(transaksi)
		}
	}
	fmt.Println("--- Histori transaksi berhasil ditampilkan ---")
}

func displayMenu(options []string) {
	fmt.Println("\nATM Menu:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}