package main

import "fmt"

type Car struct {
	unit   string
	fuelIn float32
}

func (a *Car) estimated() (int, string) {
	jarak := float32(a.fuelIn) / (1.5)
	satuan := a.unit
	return int(jarak), satuan
}
func main() {
	car1 := Car{
		fuelIn: 1.5,
		unit:   "Mill",
	}
	fmt.Printf("Bahan Bakar Terisi : %v %s\n", car1.fuelIn, "L")
	fmt.Printf("Perkiraan Jarak Tempuh : ")
	fmt.Println(car1.estimated())

	car2 := Car{
		fuelIn: 30,
		unit:   "Mill",
	}
	fmt.Printf("Bahan Bakar Terisi : %v %s\n", car2.fuelIn, "L")
	fmt.Printf("Perkiraan Jarak Tempuh : ")
	fmt.Println(car2.estimated())
}
