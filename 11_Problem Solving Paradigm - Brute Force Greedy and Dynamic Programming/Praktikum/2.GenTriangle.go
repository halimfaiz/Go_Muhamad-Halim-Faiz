package main

import "fmt"

func generateTriangle(numRow int) [][]int {
	triangle := make([][]int, numRow)
	for l := range triangle {
		triangle[l] = make([]int, l+1)
		triangle[l][0], triangle[l][l] = 1, 1
		for r := 1; r < l; r++ {
			triangle[l][r] = triangle[l-1][r-1] + triangle[l-1][r]
		}
	}
	return triangle
}
func main() {
	fmt.Println(generateTriangle(5))
}
