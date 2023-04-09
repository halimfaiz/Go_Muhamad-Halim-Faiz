package calculator

import "fmt"

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Division(a, b int) interface{} {
	if b == 0 {
		panic("Error : Cannot Divide By Zero")
	}
	result := float32(a) / float32(b)
	if result == float32(int(result)) {
		return int(result)
	} else {
		return fmt.Sprintf("%.2f", result)
	}
}

func Multiplication(a, b int) int {
	return a * b
}
