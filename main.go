package main

import "fmt"

func main() {
	var belanja int

	fmt.Println("~~~~~~~~~~Khaerul Umam~~~~~~~~~~~")
	fmt.Println("=====SELAMAT DATANG DI DUMBMERCH=====")
	fmt.Println("masukan total belanja")
	fmt.Scanf("%d", &belanja)
	if belanja <= 50000 {
		var i = belanja - (belanja * 25 / 100)
		fmt.Printf("Dumbmerch Berkah")
		fmt.Println("selamat anda mendapatkan potongan harga 25%")
		fmt.Printf("Harganya menjadi Rp %d", i)
	} else if belanja >= 70000 {
		var j = belanja - (belanja * 50 / 100)
		fmt.Println("Dumbmerch Murah Banget")
		fmt.Println("selamat anda mendapatkan potongan harga 50%")
		fmt.Printf("Harganya menjadi Rp %d", j)
	}
}
