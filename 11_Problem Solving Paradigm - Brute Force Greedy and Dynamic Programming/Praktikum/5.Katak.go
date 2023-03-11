package main

import (
	"fmt"
	"math"
)

func frog(jump []int) int {
	var cost1, cost2, jump1, jump2, min float64
	for i := 1; i < len(jump); i++ {
		jump1 = cost1 + math.Abs(float64(jump[i]-jump[i-1]))
		if i > 1 {
			jump2 = cost2 + math.Abs(float64(jump[i]-jump[i-2]))
		} else {
			jump2 = jump1
		}
		min = math.Min(jump1, jump2)
		cost2 = cost1
		cost1 = min
	}
	return int(cost1)
}
func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
