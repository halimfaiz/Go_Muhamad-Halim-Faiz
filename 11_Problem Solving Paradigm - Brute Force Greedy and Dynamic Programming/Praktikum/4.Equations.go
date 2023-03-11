package main

import "fmt"

func simpleEquations(a, b, c int) {
	for x := 1; x <= 1000; x++ {
		for y := 1; y <= 1000; y++ {
			z := a - x - y
			if x*x+y*y+z*z == c && x*y*z == b {
				fmt.Println(x, y, z) // "x = %d y = %d z = %d\n",
				return
			}
		}
	}
	fmt.Println("No Solution")
}

func main() {
	simpleEquations(1, 2, 3)  // no solution
	simpleEquations(6, 6, 14) // 1 2 3
}
