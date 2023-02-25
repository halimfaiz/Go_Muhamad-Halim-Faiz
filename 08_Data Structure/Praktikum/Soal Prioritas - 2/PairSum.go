package main

import "fmt"

func PairSum(arr []int, target int) []int {
	slicepairsum := []int{} // slice menampung hasil

	for l := 0; l < len(arr); l++ { // perulangan ke 1
		for k := 0; k < len(arr); k++ { //perulangan ke 2
			if arr[l]+arr[k] == target { //jika perulangan 1 + 2 == target akan dieksekusi
				slicepairsum = append(slicepairsum, l, k)
				return slicepairsum
			}
		}
	}
	return slicepairsum
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]

}
