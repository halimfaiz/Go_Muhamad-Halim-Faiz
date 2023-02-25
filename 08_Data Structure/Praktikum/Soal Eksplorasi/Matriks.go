package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	d1 := 0
	d2 := 0

	for kanan := 0; kanan < 3; kanan++ {
		for bawah := 0; bawah < 3; bawah++ {
			if kanan == bawah {
				d1 += matrix[kanan][bawah]
			}
			if kanan+bawah == 2 {
				d2 += matrix[kanan][bawah]
			}
		}
	}
	fmt.Println("Diagonal :", d1)
	fmt.Println("AntiDiagonal :", d2)
	fmt.Println("Perbedaan Mutlak :", int(math.Abs(float64(d1-d2))))
}
