package main

import "fmt"

func fiboTopDown(num int, lisFib map[int]int) int {
	if num <= 1 {
		return num
	}

	if val, ok := lisFib[num]; ok {
		return val
	}

	lisFib[num] = fiboTopDown(num-1, lisFib) + fiboTopDown(num-2, lisFib)
	return lisFib[num]
}
func main() {
	list := make(map[int]int)
	fmt.Println(fiboTopDown(0, list))  // 0
	fmt.Println(fiboTopDown(1, list))  // 1
	fmt.Println(fiboTopDown(2, list))  // 1
	fmt.Println(fiboTopDown(3, list))  // 2
	fmt.Println(fiboTopDown(5, list))  // 5
	fmt.Println(fiboTopDown(6, list))  // 8
	fmt.Println(fiboTopDown(7, list))  // 13
	fmt.Println(fiboTopDown(9, list))  // 34
	fmt.Println(fiboTopDown(10, list)) // 55

}
