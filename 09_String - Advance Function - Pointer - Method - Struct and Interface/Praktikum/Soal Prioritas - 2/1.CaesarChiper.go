package main

import "fmt"

func caesar(offset int, input string) string {
	out := []rune{}
	for _, caesarChiper := range input {
		caesarChiper := 'a' + (caesarChiper-'a'+rune(offset))%26
		out = append(out, caesarChiper)
	}
	return string(out)
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
