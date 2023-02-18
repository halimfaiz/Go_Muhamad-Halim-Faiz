package main

import "fmt"

func main() {
	input := 5
	for baris := 1; baris <= input; baris++ {
		for spasi := input; spasi >= baris; spasi-- {
			fmt.Printf(" ")
		}
		for bintang := 0; bintang < baris; bintang++ {
			fmt.Printf(" " + "*")
		}
		fmt.Println(" ")
	}
}
