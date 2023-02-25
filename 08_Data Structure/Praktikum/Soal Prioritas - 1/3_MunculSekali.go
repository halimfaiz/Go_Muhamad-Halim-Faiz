package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	hasil := []int{}         // slice menampung hasil
	jumlah := map[rune]int{} // map menampung element 'angka'

	for _, a := range angka { // perulangan tiap karakter string dan menghitung jumlahnya ke map 'jumlah'
		jumlah[a]++
	}
	for _, b := range angka {
		if jumlah[b] < 2 {
			int, err := strconv.Atoi(string(b)) // merubah string -> int
			if err == nil {
				hasil = append(hasil, int) // menambahkan int ke slice hasil
			}
		}
	}
	return hasil
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
