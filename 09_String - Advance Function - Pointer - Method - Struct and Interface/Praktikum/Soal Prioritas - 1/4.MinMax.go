package main

import "fmt"

func GetMinMax(numbers ...*int) (min int, max int) {
	for _, number := range numbers {
		if min > *number || min == 0 {
			min = *number
		}
		if max < *number {
			max = *number
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Input:")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Output:")
	fmt.Println(max, "is the maximum number")
	fmt.Println(min, "is the minimum number")
}
