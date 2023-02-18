package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Println("Apakah Palindrom?")
	fmt.Printf("masukan kata :")
	fmt.Scan(&text)
	fmt.Printf("Captured :")
	fmt.Println(text)

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			fmt.Println("Bukan Palindrom")
			break
		} else {
			fmt.Println("Palindrom")
			break
		}
	}
}
