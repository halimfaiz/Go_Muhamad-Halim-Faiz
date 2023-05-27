package main

import "fmt"

type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	k kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.k.kecepatanPerJam += kecepatanbaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()

	fmt.Println("kecepatan mobil cepat :", mobilcepat.k.kecepatanPerJam)
	fmt.Println("kecepatan mobil lambat :", mobillamban.k.kecepatanPerJam)
}
