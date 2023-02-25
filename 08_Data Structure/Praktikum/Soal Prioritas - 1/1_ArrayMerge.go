package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	data := make(map[string]bool) // map menampung data arrayA & arrayB
	merge := []string{}           // slice untuk menampung hasil

	for _, value := range arrayA {
		if _, ada := data[value]; ada == false {
			data[value] = true
			merge = append(merge, value)
		}
	}
	for _, value := range arrayB {
		if _, ada := data[value]; ada == false {
			merge = append(merge, value)
		}
	}
	return merge
}

func main() {
	// Test Case
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
