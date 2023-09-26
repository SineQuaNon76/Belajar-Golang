package main

import "fmt"

func main(){
	fmt.Println("Hello Guys")
	fmt.Println("Persija" + " " + "String")
	fmt.Println(false) // false
	fmt.Println(false || true) //true


	var pemain string = "Firza"
	println(pemain)


	name := "Bias"
	angka := 12

	
	//Multi Variable
	huruf, angka, nama := "Abc", 12, "Goday"
	fmt.Printf("Ini Multi Variable loh %s %d %s\n", huruf, angka, nama)


	//_UnderScore
	_ = "Apa Ya"


	println(angka)
	println(name)

	fmt.Println("hallo" + pemain + name)
	fmt.Printf("halo %s %s!\n", pemain, name)
}
