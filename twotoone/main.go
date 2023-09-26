package main

import "fmt"

func main() {
	s1 := "abcde"
	s2 := "bcdefg"
	result := TwoToOne(s1, s2)
	fmt.Println(result)
}

func TwoToOne(s1 string, s2 string) string {
	tampung := make(map[rune]bool)

	for _, kata := range s1 {
		tampung[kata] = true
	}

	for _, kata := range s2 {
		tampung[kata] = true
	}

	hasil := ""
	for kata := range tampung {
		hasil += string(kata)
	}

	return hasil
}
