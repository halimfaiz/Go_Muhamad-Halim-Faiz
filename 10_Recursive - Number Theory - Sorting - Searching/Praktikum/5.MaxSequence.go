package main

import "fmt"

func MaxSequence(arr []int) int {
	maxNow := 0
	maxSequ := 0
	for _, val := range arr {
		maxNow += val
		if maxNow < 0 {
			maxNow = 0
		} else if maxNow > maxSequ {
			maxSequ = maxNow
		}
	}
	return maxSequ
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
