package main

import (
	"fmt"
)

func isPrime(n, num int) bool {
	if num == 1 {
		return true
	}
	if n%num == 0 {
		return false
	}
	return isPrime(n, num-1)
}

func getNthPrime(n, num int) int {
	if isPrime(num, num-1) {
		if n == 1 {
			return num
		}
		return getNthPrime(n-1, num+1)
	}
	return getNthPrime(n, num+1)

}

func primeX(n int) int {
	return getNthPrime(n, 2)
}

func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
