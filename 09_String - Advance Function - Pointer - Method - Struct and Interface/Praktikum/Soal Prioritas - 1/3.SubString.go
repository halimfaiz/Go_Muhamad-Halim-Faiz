package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var hasil string
	awal := strings.Contains(b, a)
	if awal == true {
		hasil += a
	}
	akhir := strings.Contains(a, b)
	if akhir == true {
		hasil += b
	}
	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
