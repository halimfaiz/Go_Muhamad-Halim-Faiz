package main

import "fmt"

func angkaBiner(number int) {
	for b := 0; b <= number; b++ {
		fmt.Printf("%b ", b)
	}
}

func main() {
	angkaBiner(2)
	fmt.Println()
	angkaBiner(3)
}
