package main

import "fmt"

func romawi(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	convert := ""
	for r := 0; r < len(val); r++ {
		for num >= val[r] {
			convert += romawi[r]
			num -= val[r]
		}
	}
	return convert
}

func main() {
	fmt.Println(romawi(4))    // IV
	fmt.Println(romawi(9))    // IX
	fmt.Println(romawi(23))   //XXIII
	fmt.Println(romawi(2021)) //MMXXI
	fmt.Println(romawi(1646)) //MDCXLVI
}
