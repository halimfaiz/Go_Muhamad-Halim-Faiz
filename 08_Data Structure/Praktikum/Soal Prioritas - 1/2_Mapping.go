package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// map 'total' untuk menampung key dan jumlah value
	total := map[string]int{}
	// perulangan untuk mengisi key dan jumlah value
	for _, v := range slice {
		total[v] += 1
	}
	return total
}
func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}
