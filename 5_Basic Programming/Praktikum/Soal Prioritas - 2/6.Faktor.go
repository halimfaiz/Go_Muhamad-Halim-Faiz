package main

import "fmt"

func main() {
	number := 26
	fmt.Println("angka : ", number)
	for i := 1; i <= number; i++ {
		factorial := number % i
		if factorial == 0 {
			fmt.Println(i)
		}
	}
}
