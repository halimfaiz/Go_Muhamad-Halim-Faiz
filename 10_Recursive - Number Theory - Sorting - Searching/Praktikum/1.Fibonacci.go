package main

import "fmt"

func fibonacci(number int) int {
	fibo := 0
	if number == 0|1 {
		return number
	}
	if number > 1 {
		fibo = fibonacci(number-1) + fibonacci(number-2)
	}
	return fibo
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
