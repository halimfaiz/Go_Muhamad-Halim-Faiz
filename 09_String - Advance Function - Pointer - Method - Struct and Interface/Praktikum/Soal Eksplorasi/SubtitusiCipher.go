package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var encoded = ""
	s.nameEncode += ChiperSubt(s.name)
	encoded += ChiperSubt(s.name)
	return encoded
}

func (s *student) Decode() string {
	var decoded = ""
	if s.nameEncode == "" {
		s.Encode()
	}
	decoded += ChiperSubt(s.nameEncode)

	return decoded
}

func ChiperSubt(s string) string {
	var hasil string
	for _, word := range s {
		encodedWord := 'z' - rune(word) + 'a'
		hasil += string(encodedWord)
	}
	return hasil
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of students name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of students name " + a.name + " is : " + c.Decode())
	}
}
